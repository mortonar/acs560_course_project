package handlers

import (
    "github.com/mortonar/acs560_course_project/messaging/messages/request"
    "github.com/mortonar/acs560_course_project/messaging/messages/response"
    "github.com/mortonar/acs560_course_project/database/models"
    "github.com/jinzhu/gorm"
    "fmt"
)

// TODO error handling
func HandleAddBook(addReq request.AddBook, db *gorm.DB, userId uint) (*response.Base, error) {
    fmt.Println("user ID : ", userId)
    shelf := models.Shelf{}
    db.First(&shelf, "user_id = ? AND name = ?", userId, addReq.ShelfName)
    fmt.Println("shelf: ", shelf)

    book := models.Book{}
    db.FirstOrCreate(&book, addReq.Book)
    fmt.Println("book: ", shelf)

    db.Model(&shelf).Association("Books").Append(book)
    db.Update(&shelf)

    return &response.Base{
        Success: true,
        Status: "Successfully added book",
        Payload: nil,
    }, nil
}
