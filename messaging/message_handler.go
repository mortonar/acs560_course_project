package messaging

import (
    "fmt"
    "github.com/mortonar/acs560_course_project/messaging/messages/request"
    "github.com/mortonar/acs560_course_project/messaging/messages/response"
    "github.com/mortonar/acs560_course_project/messaging/handlers"
    "github.com/mortonar/acs560_course_project/database"
    "github.com/mortonar/acs560_course_project/database/models"
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
            handler.responseChan <- response.Base{true, "Got CreateAccount Message: " + message.Token, nil}
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
            payload := response.Login{ Token: handler.session.Token }
            handler.responseChan <- response.Base{true, "Got Auth Message", payload }
	    case "Search":
            var bookSearch = request.BookSearch{}
            error := ParseMessage(message, &bookSearch)
            if error == nil {
                searchResp, err := handlers.HandleBookSearch(bookSearch)
                if err == nil {
                    baseResponse := response.Base{Success:true, Status: "Successful Search", Payload: *searchResp}
                    handler.responseChan <- baseResponse
                } else {
                    fmt.Println("Error in search: ", err)
                    baseResponse := response.Base{
                        Success:false,
                        Status: fmt.Sprintf("Error in search: %s", err),
                        Payload: nil,
                    }
                    handler.responseChan <- baseResponse
                }
            } else {
                baseResponse := response.Base{
                    Success:false,
                    Status: fmt.Sprintf("Error in parsing message: %s", error),
                    Payload: nil,
                }
                handler.responseChan <- baseResponse
			}
        }
        // TODO ensure session exists before allowing other actions
    }
}
