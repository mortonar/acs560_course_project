package handlers

import (
    "github.com/mortonar/acs560_course_project/messaging/messages/request"
    "github.com/mortonar/acs560_course_project/messaging/messages/response"
    "github.com/mortonar/acs560_course_project/database/models"
    "github.com/jinzhu/gorm"
)

// TODO error handling
func HandleAddBook(addReq request.AddBook, db *gorm.DB, user models.User) (*response.Base, error) {
    shelf := user.GetShelf(addReq.ShelfName)
    shelf.Books = append(shelf.Books, addReq.Book)
    db.Save(&shelf)
    return &response.Base{
        Success: true,
        Status: "Successfully added book",
        Payload: nil,
    }, nil
}
