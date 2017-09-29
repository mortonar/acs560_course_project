package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/mortonar/acs560_course_project/models"
)

func main() {
    args := "user=booktracker dbname=booktracker sslmode=disable password=booktracker"
    db, err := gorm.Open("postgres", args)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    db.AutoMigrate(&models.Book{})
    db.AutoMigrate(&models.Shelf{})
    db.AutoMigrate(&models.User{})
}