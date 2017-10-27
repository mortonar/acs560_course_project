package models

import "github.com/jinzhu/gorm"

type Session struct {
    gorm.Model
    UserID uint `gorm:"primary_key"`
    Token string `gorm:"primary_key"`
}

