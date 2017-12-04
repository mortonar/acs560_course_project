package handlers

import (
    "github.com/mortonar/acs560_course_project/messaging/messages/request"
    "fmt"
    "github.com/mortonar/acs560_course_project/database/models"
    "github.com/jinzhu/gorm"
    "github.com/satori/go.uuid"
    "errors"
)

func HandleCreateAccount(request request.CreateAccount, connection *gorm.DB) {
    fmt.Println(request)
    user := models.User{Login: request.UserName, Email: request.Email, Password: request.Password}

    connection.Create(&user)
    models.CreateShelvesForUser(connection, user.ID)
}

func HandleLogin(request request.AuthRequest, db *gorm.DB) (*models.Session, error) {
    fmt.Println("Handling login...")
    user := models.User{}
    fmt.Println("Looking for user : ", request.UserName, " | ", request.EncryptedPass)
    db.Where("login = ? AND password = ?", request.UserName, request.EncryptedPass).First(&user)
    fmt.Println("Found user : ", user.Login, " | ", user.Password, " | ", user.ID)
    if &user != nil {
        token := uuid.NewV4().String()
        fmt.Println("Generating token for user : ", token)
        session := models.Session{UserID: user.ID, Token: token}
        db.Create(&session)
        fmt.Println("Generating new session for user : ", session)
        return &session, nil
    } else {
        return nil, errors.New(fmt.Sprint("Cannot find user with login %v", request.UserName))
    }
}
