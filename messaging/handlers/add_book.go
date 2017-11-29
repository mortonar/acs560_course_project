package handlers

import (
    "github.com/mortonar/acs560_course_project/messaging/messages/request"
    "github.com/mortonar/acs560_course_project/messaging/messages/response"
    "github.com/mortonar/acs560_course_project/database/models"
    "github.com/jinzhu/gorm"
)

// TODO error handling
func HandleAddBook(addReq request.AddBook, db *gorm.DB, user models.User) (*response.Base, error) {
    shelf := models.Shelf{}
    db.First(&shelf, "user_id = ? AND name = ?", user.ID, addReq.ShelfName)

    book := models.Book{}
    db.FirstOrCreate(&book, addReq.Book)

    db.Model(&shelf).Association("Books").Append(book)
    db.Update(&shelf)

    return &response.Base{
        Success: true,
        Status: "Successfully added book",
        Payload: nil,
    }, nil
}
