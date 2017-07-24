package main

import (
    "os"
    "os/signal"
    "fmt"
    "log"
    "time"
    "syscall"

    "polyglotenv"
)

type GracefullyEnds interface {
    GracefullyShutdown()
}

func polyglot() {
    var controller *polyglotenv.Controller

    ch := make(chan os.Signal)
    signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

    polyglotSocket := "/tmp/polyglot.sock"

    if _, err := os.Stat(polyglotSocket); os.IsNotExist(err) {
        fmt.Println("Socket doesn't exist yet")
        logger := log.New(os.Stdout, "server: ", log.Lshortfile)
        
        controller = polyglotenv.NewController(logger)
        go controller.Listen(polyglotSocket)
    }

    time.Sleep(100 * time.Millisecond)

    logger := log.New(os.Stdout, "client 1: ", log.Lshortfile)
    client1 := &polyglotenv.PolyglotClient{logger}
    client1.Connect(polyglotSocket)

    <-ch
    fmt.Println("Shutting down gracefully...")
    controller.GracefullyShutdown()
    time.Sleep(1 * time.Second)
    os.Exit(0)
}




func main() {
    fmt.Println("polyglot by @liamzebedee")
    fmt.Println("Support developer experience - liamz.co")

    if len(os.Args) < 2 {
        return
    }
    firstArg := os.Args[1]
    if len(firstArg) > 0 && firstArg[0] == '-' {
        fmt.Println("It is a command")
    } else {
        polyglot()

        fmt.Println("Let's build")
    }
}