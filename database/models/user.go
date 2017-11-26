package models

import "github.com/jinzhu/gorm"

type User struct {
    gorm.Model
    Login string `gorm:"primary_key"`
    Email string `gorm:"primary_key"`
    Password string
    Shelves []Shelf
    Session Session
}

func (user *User) GetShelf(name string) *Shelf {
    shelf := Shelf{}
    for _, s := range user.Shelves {
       if s.Name == name {
           shelf = s
           break
       }
    }
    return &shelf
}
