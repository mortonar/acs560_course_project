package sever

import (
    "net"
    "log"
)

// TODO possible include a channel between the main server thread and the server start so we can send it commands?
type Server struct {
    Port int
    clientConnections []*ClientConnection
}

func (server *Server) Start() {
    listener, error := net.Listen("tcp", ":8000")
    defer listener.Close()
    if error != nil {
        panic(error)
    }

    for {
        client, error := listener.Accept()
        defer client.Close()
        if error != nil {
            log.Fatalln(error)
            continue
        }
        cl := NewClientConnection(client)
        server.clientConnections = append(server.clientConnections, cl)
        cl.Start()
    }
}
