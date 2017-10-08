package messaging

import (
    "fmt"
    "github.com/mortonar/acs560_course_project/messaging/messages/request"
    "github.com/mortonar/acs560_course_project/messaging/messages/response"
    "github.com/mortonar/acs560_course_project/messaging/handlers"
    "encoding/json"
)

type MessageHandler struct {
    requestChan <-chan request.Base
    responseChan chan<- response.Base
}

func NewMessageHandler(requestChan <-chan request.Base, responseChan chan<- response.Base) *MessageHandler {
    mh := &MessageHandler{requestChan, responseChan}
    return mh
}

// TODO one router per client connection or one master router?
func (handler *MessageHandler) Start() {
    go handler.process()
}

func (handler *MessageHandler) Stop() {
}

func (handler *MessageHandler) process() {
    for {
        message := <-handler.requestChan
        fmt.Println("MessageHandler::gotMessage ->")
        fmt.Println(message)
        switch message.Action {
        case "Auth":
            bytes, _ := json.Marshal(message.Payload) // TODO error handling
            var authReq = request.AuthRequest{}
            _ = json.Unmarshal(bytes, authReq)
            handlers.HandleLogin(authReq)
        }
        handler.responseChan <- response.Base{true, "Got Message: " + message.Token}
    }
}
