package main

import (
    "os"
    "fmt"
    "net"
)



func polyglot() {
    polyglotSocket := "/tmp/polyglot.sock"
    addr := &net.UnixAddr{polyglotSocket, "unix"}


    needToStartManager := false
    if _, err := os.Stat(polyglotSocket); os.IsNotExist(err) {
        panic("Socket doesn't exist yet")
        list, err := net.ListenUnix("unix", addr)

        go func() {
            conn, err := list.AcceptUnix()
            if err != nil {
                panic("Couldn't accept conn :"+err.Error())
            }

            go func() {
                buf := make([]byte, 2048)
                for {
                    buf, err := conn.ReadFromUnix()
                }
            }
        }

        list.SetUnlinkOnClose(true)
    }

    conn, err := net.DialUnix("unix", nil, addr)
    if err != nil {
        panic("Error connecting to Polyglot controller")
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