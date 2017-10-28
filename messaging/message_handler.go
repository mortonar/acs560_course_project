package messaging

import (
    "fmt"
    "github.com/mortonar/acs560_course_project/messaging/messages/request"
    "github.com/mortonar/acs560_course_project/messaging/messages/response"
    "github.com/mortonar/acs560_course_project/messaging/handlers"
    "github.com/mortonar/acs560_course_project/database"
    "github.com/mortonar/acs560_course_project/database/models"
    "encoding/json"
)

type MessageHandler struct {
    requestChan <-chan request.Base
    responseChan chan<- response.Base
    dbProxy *database.DBProxy
    session *models.Session
}

func NewMessageHandler(requestChan <-chan request.Base, responseChan chan<- response.Base) *MessageHandler {
    mh := &MessageHandler{requestChan, responseChan, database.NewDBProxy(), nil}
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
        case "CreateAccount":
            fmt.Println("Creating account...")
            var createReq = request.CreateAccount{}
            error := ParseMessage(message, &createReq)
            if error == nil {
                handlers.HandleCreateAccount(createReq, handler.dbProxy.GetConnection())
            } else {
                fmt.Println("Error: ", error)
            }
        case "Auth":
            var authReq = request.AuthRequest{}
            error := ParseMessage(message, &authReq)
            if error == nil {
                session, err := handlers.HandleLogin(authReq, handler.dbProxy.GetConnection())
                if err == nil {
                    handler.session = session
                } else {
                    fmt.Println(err)
                }

            }
	    case "Search":
            bytes, err := json.Marshal(message.Payload) // TODO actual error handling
            if err != nil {
                fmt.Println("Marshal ERROR: ", err)
            }
            decoded := string(bytes)
            fmt.Printf("DECODED: %+v (%T)\n", decoded)
		    // TODO - perform search and  format and send back response
        }		
        // TODO ensure session exists before allowing other actions
        handler.responseChan <- response.Base{true, "Got Message: " + message.Token}
    }
}
