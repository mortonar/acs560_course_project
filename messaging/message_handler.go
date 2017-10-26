package messaging

import (
    "fmt"
    "github.com/mortonar/acs560_course_project/messaging/messages/request"
    "github.com/mortonar/acs560_course_project/messaging/messages/response"
    "github.com/mortonar/acs560_course_project/messaging/handlers"
    "github.com/mortonar/acs560_course_project/database"
)

type MessageHandler struct {
    requestChan <-chan request.Base
    responseChan chan<- response.Base
    dbProxy *database.DBProxy
}

func NewMessageHandler(requestChan <-chan request.Base, responseChan chan<- response.Base) *MessageHandler {
    mh := &MessageHandler{requestChan, responseChan, database.NewDBProxy()}
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
        fmt.Println("MessageHandler::gotMessage ->\n%v", message)
        switch message.Action {
        case "Auth":
            var authReq = request.AuthRequest{}
            error := ParseMessage(message, &authReq)
            if error == nil {
                handlers.HandleLogin(authReq)
            }
        case "CreateAccount":
            fmt.Println("Creating account...")
            var createReq = request.CreateAccount{}
            error := ParseMessage(message, &createReq)
            if error == nil {
                handlers.HandleCreateAccount(createReq, handler.dbProxy.GetConnection())
            } else {
                fmt.Println("Error: ", error)
            }
        }
        handler.responseChan <- response.Base{true, "Got Message: " + message.Token}
    }
}
