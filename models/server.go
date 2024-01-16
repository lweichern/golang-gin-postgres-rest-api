package models

type Book struct { // start with capital case to make it public for other external modules to use, 
	ID 		string	`json:"id"` // json to serialize the field so that the server can convert data to real json object
	Title	string	`json:"title"`
	Author 	string	`json:"author"`
	Quantity int16 `json:"quantity"`
}