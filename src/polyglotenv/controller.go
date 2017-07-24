package polyglotenv

import (
    "log"
    "net"
)




type Controller struct {
    log *log.Logger
    shutdown chan bool
}

func NewController(log *log.Logger) *Controller {
    controller := Controller{}
    controller.log = log
    controller.shutdown = make(chan bool)
    return &controller
}

func (s *Controller) GracefullyShutdown() {
    s.shutdown <- true
}

func (s *Controller) Listen(addr string) {
    list, err := net.Listen("unix", addr)
    defer list.Close()
    defer func(){
        // if err := os.Remove("/tmp/polyglot.sock"); err != nil {
            // panic(err)
        // }
    }()

    newConnections := make(chan *net.Conn)
    go func(){
        for {
            conn, err := list.Accept()
            if err != nil {
                // panic("Couldn't accept conn - " + err.Error())
            }
            
            newConnections <- &conn
        }
    }()

    if err != nil {
        s.log.Println("Listening on Unix socket failed")
        panic(err)
    }
    defer list.Close()

    for {
        select {
        case <-s.shutdown:
            return
        case conn := <-newConnections:
            go s.handleConnection(conn)
        }
    }
}

func (s *Controller) handleConnection(conn *net.Conn) {
    return
}