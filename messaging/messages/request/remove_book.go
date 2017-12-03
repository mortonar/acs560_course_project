package request

import "github.com/mortonar/acs560_course_project/database/models"

type RemoveBook struct {
    Book models.Book
    ShelfName string
}
