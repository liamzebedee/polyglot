package main

import (
    "os"
    "fmt"
)

func main() {
    fmt.Println("polyglot by @liamzebedee")
    fmt.Println("Support developer experience - liamz.co")

    if len(os.Args) < 2 {
        return
    }

    firstArg := os.Args[1]
    if firstArg[0] == '-' {
        fmt.Println("It is a command")
    } else {
        fmt.Println("Let's build")
    }
}