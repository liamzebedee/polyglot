package main

import (
    "os"
    "fmt"
    "net"
    "log"
    "time"
)

// type Client interface{}
type PolyglotServer struct {
    log *log.Logger
//     clients map[string]*Client
}

func (s *PolyglotServer) Listen(addr *net.UnixAddr) {
    list, err := net.ListenUnix("unix", addr)
    if err != nil {
        s.log.Println("FUCK LISTEN FAILED")
        panic(err)
    }
    defer list.Close()

    go func() {
        _, err = list.AcceptUnix()
        if err != nil {
            panic("Couldn't accept conn :"+err.Error())
        }

        s.log.Println("New client!")

        // go func() {
        //     buf := make([]byte, 2048)
        //     for {
        //         buf, err := conn.ReadFromUnix()
        //     }

        // }
    }()

    // list.SetUnlinkOnClose(true)
}


func polyglot() {
    polyglotSocket := "/tmp/polyglot.sock"
    addr := &net.UnixAddr{polyglotSocket, "unix"}


    if _, err := os.Stat(polyglotSocket); os.IsNotExist(err) {
        fmt.Println("Socket doesn't exist yet")
        logger := log.New(os.Stdout, "server: ", log.Lshortfile)
        server := &PolyglotServer{logger}
        go server.Listen(addr)
        time.Sleep(100 * time.Millisecond)
    }

    _, err := net.DialUnix("unix", nil, addr)
    if err != nil {

        panic("Error connecting to Polyglot controller :"+err.Error())
    }
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