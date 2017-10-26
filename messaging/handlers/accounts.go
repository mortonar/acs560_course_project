package handlers

import (
    "github.com/mortonar/acs560_course_project/messaging/messages/request"
    "fmt"
    "github.com/mortonar/acs560_course_project/database/models"
    "github.com/jinzhu/gorm"
)

func HandleCreateAccount(request request.CreateAccount, connection *gorm.DB) {
    fmt.Println(request)
    user := models.User{Login: request.UserName, Email: request.Email, Password: request.Password}

    connection.Create(&user)
}

func HandleLogin(request request.AuthRequest) {
    fmt.Println("Login message: ", request)
    fmt.Println("Logging in...")
    fmt.Println("Logged in...")
}
