package polyglotenv

import (
    "log"
    "net"
)

type PolyglotClient struct {
    Log *log.Logger
}

func (s *PolyglotClient) Connect(addr string)  {
    // clientAddr := &net.UnixAddr{"/tmp/whatever", "unix"}
    conn, err := net.Dial("unix", addr)
    if err != nil {
        panic("Error connecting to Polyglot controller - "+err.Error())
    }

    conn.Write([]byte{0,1,2,3})
    conn.Close()
}