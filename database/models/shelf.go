package models

import "github.com/jinzhu/gorm"

type Shelf struct {
    gorm.Model
    Name string `gorm:"primary_key"`
    Books []Book `gorm:"many2many:shelves_books;"`
    UserId uint `gorm:"primary_key"`
}
