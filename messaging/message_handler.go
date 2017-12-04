package messaging

import (
    "fmt"
    "github.com/mortonar/acs560_course_project/messaging/messages/request"
    "github.com/mortonar/acs560_course_project/messaging/messages/response"
    "github.com/mortonar/acs560_course_project/messaging/handlers"
    "github.com/mortonar/acs560_course_project/database"
    "github.com/mortonar/acs560_course_project/database/models"
    "errors"
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


var loginExemptActions = []string{"CreateAccount", "Auth"}

func actionRequiresLogin(action string) bool {
    for _, a := range loginExemptActions {
        if a == action {
            return false
        }
    }
    return true
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

        if actionRequiresLogin(message.Action) && handler.session == nil {
            reason := "User must be logged in to perform " + message.Action
            handler.handleGenericError(reason, errors.New(reason))
            continue
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
                searchResp, err := handlers.HandleBookList(bookList, handler.dbProxy.GetConnection(), handler.session.UserID)
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
                addResp, err := handlers.HandleAddBook(addBook, handler.dbProxy.GetConnection(), handler.session.UserID)
                if err == nil {
                    handler.responseChan <- *addResp
                } else {
                    handler.handleGenericError("Error in adding book to list: %s", err)
                }
            }
        case "RemoveBook":
            fmt.Println("Got an RemoveBook Request")
            var removeBook = request.RemoveBook{}
            if handler.parse(message, &removeBook) {
                removeResp, err := handlers.HandleRemoveBook(removeBook, handler.dbProxy.GetConnection(), handler.session.UserID)
                if err == nil {
                    handler.responseChan <- *removeResp
                } else {
                    handler.handleGenericError("Error in removing the book from the list: %s", err)
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
