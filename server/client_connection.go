package server

import (
    "net"
    "encoding/json"
    "github.com/mortonar/acs560_course_project/messaging/messages/request"
    "github.com/mortonar/acs560_course_project/messaging"
    "github.com/mortonar/acs560_course_project/messaging/messages/response"
    "fmt"
)

// TODO what do we do when the client closes their connection?
type ClientConnection struct {
    Connection net.Conn
    requestChan chan request.Base
    responseChan chan response.Base
    handler *messaging.MessageHandler
}

func NewClientConnection(connection net.Conn) *ClientConnection {
    requestChan := make(chan request.Base)
    responseChan := make(chan response.Base)
    handler := messaging.NewMessageHandler(requestChan, responseChan)
    cc := &ClientConnection {
        connection,
        requestChan,
        responseChan,
        handler,
    }
    return cc
}

// TODO start vs run?
func (clientConn *ClientConnection) Start() {
    clientConn.handler.Start()
    go clientConn.handleRequests()
    go clientConn.handleResponses()
}

func (clientConn *ClientConnection) Stop() {
    clientConn.handler.Stop()
    close(clientConn.requestChan)
    close(clientConn.responseChan)
}

func (clientConn *ClientConnection) handleRequests() {
    empty := request.Base{}
    for {
        fmt.Println("waiting for message...")
        jsonDecoder := json.NewDecoder(clientConn.Connection)
        fmt.Println("received message!")
        var message request.Base
        jsonDecoder.Decode(&message)
        if message == empty {
            fmt.Println("empty message! stopping clientconnection!")
            clientConn.Stop()
            break
        }
        fmt.Println(message)
        clientConn.requestChan <- message
    }
}

func (clientConn *ClientConnection) handleResponses() {
    for resp := range clientConn.responseChan {
        bytes, err := json.Marshal(resp)
        if err != nil {
            fmt.Println(err)
        } else {
            clientConn.Connection.Write(bytes)
        }
    }
}
