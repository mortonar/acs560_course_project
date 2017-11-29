package models

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model   `json:"-"` // ignore gorm fields when marshaling
	ISBN13       string
	Title        string
	Author       string
	Description  string
	ThumbnailURL string
	ImageURL     string
	Shelves      []Shelf `gorm:"many2many:shelves_books;"`
}
