package models

import (
    "github.com/jinzhu/gorm"
)

type Shelf struct {
    gorm.Model
    Name string `gorm:"primary_key"`
    Books []Book `gorm:"many2many:shelves_books;"`
    UserId uint `gorm:"primary_key"`
}

func CreateShelvesForUser(conn *gorm.DB, userId uint) {
    shelves := [3]string{"Read", "Reading", "To Read"}
    for _, shelfName := range shelves {
        shelf := Shelf {
            Name: shelfName,
            UserId: userId,
        }
        conn.FirstOrCreate(&shelf, shelf)
    }
}
