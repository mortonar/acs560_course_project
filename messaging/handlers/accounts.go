package handlers

import (
    "github.com/mortonar/acs560_course_project/messaging/messages/request"
    "fmt"
)

func HandleLogin(request request.AuthRequest) {
    fmt.Println("Login message: ", request)
    fmt.Println("Logging in...")
    fmt.Println("Logged in...")
}
