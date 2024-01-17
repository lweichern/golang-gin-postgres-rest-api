package models

import (
	"gorm.io/gorm"
)

type Book struct { // start with capital case to make it public for other external modules to use, 
	gorm.Model
	Title	string	`json:"title"`
	Author 	string	`json:"author"`
	Quantity int16 `json:"quantity"`
}