package response

import "github.com/mortonar/acs560_course_project/database/models"

type BookList struct {
    Books []models.Book
}
