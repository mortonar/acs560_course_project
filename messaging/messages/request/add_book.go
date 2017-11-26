package request

import "github.com/mortonar/acs560_course_project/database/models"

type AddBook struct {
    Book models.Book
    ShelfName string
}
