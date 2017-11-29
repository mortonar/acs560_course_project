package handlers

import (
    "github.com/jinzhu/gorm"
    "github.com/mortonar/acs560_course_project/database/models"
    "github.com/mortonar/acs560_course_project/messaging/messages/request"
    "github.com/mortonar/acs560_course_project/messaging/messages/response"
    "errors"
    "fmt"
)

// Handles a Book List request, returning a response containing the book list
func HandleBookList(searchReq request.BookList, db *gorm.DB) (*response.BookList, error) {
    fmt.Println("Handling book list request...")
    Shelf := models.Shelf{}
    name := searchReq.Name
    db.Preload("Books").First(&Shelf, "name = ?", name)
    fmt.Println("Found BookList named " + name)
    
    if &Shelf != nil {
        fmt.Println(len(Shelf.Books))
        response := response.BookList{Books: Shelf.Books}
        return &response, nil
    } else {
        return nil, errors.New(fmt.Sprint("Cannot find book list with name %v", name))
    }
}
