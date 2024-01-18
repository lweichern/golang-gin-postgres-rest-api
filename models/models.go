package models

import (
	"gorm.io/gorm"
)

type Book struct { // start with capital case to make it public for other external modules to use, 
	gorm.Model
	Title	string	`json:"title"`
	Author 	string	`json:"author"`
	Quantity int16 `json:"quantity"`
	AuthorID uint	`json:"authorId"`
}

type User struct {
	gorm.Model
	Username 	string `json:"username"`
	Password 	string `json:"password"`
}

type Author struct {
	gorm.Model 
	Name 	string `json:"name"`
	Country	string `json:"country"`
	Age		int16  `json:"age"`
	Books	[]Book
}