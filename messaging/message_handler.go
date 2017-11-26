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
    empty := request.Base{}
    for {
        message := <-handler.requestChan
        fmt.Println("MessageHandler::gotMessage ->\n%v", message)
        if message == empty {
            fmt.Println("empty message! stopping messagehandler!")
            break
        }
        switch message.Action {
        case "CreateAccount":
            fmt.Println("Creating account...")
            var createReq = request.CreateAccount{}
            if handler.parse(message, &createReq) {
                handlers.HandleCreateAccount(createReq, handler.dbProxy.GetConnection())
                handler.responseChan <- response.Base{true, "Got CreateAccount Message: " + message.Token, nil}
            }
        case "Auth":
            var authReq = request.AuthRequest{}
            if handler.parse(message, &authReq) {
                session, err := handlers.HandleLogin(authReq, handler.dbProxy.GetConnection())
                if err == nil {
                    handler.session = session
                    payload := response.Login{ Token: handler.session.Token }
                    handler.responseChan <- response.Base{true, "Login Success", payload }
                } else {
                    handler.handleGenericError("Login error", err)
                }
            }
        // TODO ensure session exists before processing actions
        case "BookSearch":
            var bookSearch = request.BookSearch{}
            if handler.parse(message, &bookSearch) {
                searchResp, err := handlers.HandleBookSearch(bookSearch)
                if err == nil {
                    baseResponse := response.Base{Success:true, Status: "Successful Search", Payload: *searchResp}
                    handler.responseChan <- baseResponse
                } else {
                    handler.handleGenericError("Error in search: %s", err)
                }
            }
        case "BookList":
            fmt.Println("Got a BookList Request")
            var bookList = request.BookList{}
            if handler.parse(message, &bookList) {
                fmt.Println("Making HandleBookListRequest")
                searchResp, err := handlers.HandleBookList(bookList, handler.dbProxy.GetConnection())
                if err == nil {
                    baseResponse := response.Base{Success:true, Status: "Successful Book List Request", Payload: *searchResp}
                    handler.responseChan <- baseResponse
                } else {
                    handler.handleGenericError("Error in requesting book list: %s", err)
                }
            }
        case "AddBook":
            fmt.Println("Got an AddBook Request")
            var addBook = request.AddBook{}
            if handler.parse(message, &addBook) {
                user := handler.dbProxy.GetBookTrackerUser()
                addResp, err := handlers.HandleAddBook(addBook, handler.dbProxy.GetConnection(), *user)
                if err == nil {
                    handler.responseChan <- *addResp
                } else {
                    handler.handleGenericError("Error in adding book to list: %s", err)
                }
            }
        }
    }
}

func (handler *MessageHandler) parse(baseMessage request.Base, intendedType interface{}) bool {
    error := ParseMessage(baseMessage, intendedType)
    if error == nil {
        return true
    } else {
        baseResponse := response.Base{
            Success:false,
            Status: fmt.Sprintf("Error in parsing message: %s", error),
            Payload: nil,
        }
        handler.responseChan <- baseResponse

        return false
    }
}

func (handler *MessageHandler) handleGenericError(message string, err error) {
    fmt.Println(message + ":", err)
    baseResponse := response.Base{
        Success:false,
        Status: fmt.Sprintf("Error in adding book to list: %s", err),
        Payload: nil,
    }
    handler.responseChan <- baseResponse
}
