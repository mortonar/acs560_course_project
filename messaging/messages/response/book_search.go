package response

import "github.com/mortonar/acs560_course_project/database/models"

type BookSearch struct {
    Books []models.Book
}
