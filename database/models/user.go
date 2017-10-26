package models

import "github.com/jinzhu/gorm"

type User struct {
    gorm.Model
    Login string `gorm:"primary_key"`
    Email string `gorm:"primary_key"`
    Password string
    Shelves []Shelf
}
