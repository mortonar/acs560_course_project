package handlers

import (
    "fmt"
    "github.com/mortonar/acs560_course_project/messaging/messages/request"
    "github.com/mortonar/acs560_course_project/messaging/messages/response"
    "github.com/mortonar/acs560_course_project/database/models"
    "github.com/jinzhu/gorm"
)

// TODO error handling
func HandleRemoveBook(removeReq request.RemoveBook, db *gorm.DB, userId uint) (*response.Base, error) {
    fmt.Println("Removing the Book " + removeReq.Book.Title + " from the list " + removeReq.ShelfName)
    shelf := models.Shelf{}
    db.First(&shelf, "user_id = ? AND name = ?", userId, removeReq.ShelfName)

    book := models.Book{}
    db.First(&book, removeReq.Book)

    db.Exec("DELETE FROM shelves_books  WHERE (shelf_id = ?) AND (shelf_name = ?) AND (shelf_user_id = ?) AND (book_id IN (?))", shelf.ID, shelf.Name, userId, book.ID)

    return &response.Base{
        Success: true,
        Status: "Successfully removed the book from the shelf",
        Payload: nil,
    }, nil
}
