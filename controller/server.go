package controller

import (
	"encoding/json"
	"errors"
	"example/http-server/database"
	"example/http-server/models"
	"fmt"
	"net/http"

	// "crypto/rand"
	// "encoding/hex"
	// "encoding/json"

	"github.com/gin-gonic/gin"
)

// tokens slice to store token generated for verfication purpose
var tokens []string

func getBookById(id string) (*models.Book, error) {
	bookList, err := database.Db.Query("SELECT * FROM books WHERE id=$1", id)
	
	if err != nil {
		return nil, errors.New("Something went wrong")
	}

	for bookList.Next() {
		var b models.Book
		if err := bookList.Scan(&b.ID, &b.Title, &b.Author, &b.Quantity); err != nil {
			fmt.Println("error: ", err)
			return nil, errors.New("Book not found")
		}
		return &b, nil
	}

	return nil, errors.New("Something went wrong")
	// // loop through books array and find the book that matches with the ID
	// for index, book := range books {
	// 	if book.ID == id {
	// 		return &books[index], nil
	// 	}
	// }

	// // return error message if book is not found
	// return nil, errors.New("Book not found")
}

func GetBooks(c *gin.Context){
	// SELECT query
	bookList, err := database.Db.Query("SELECT * FROM books")

	// Check errors
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(400, "Something went wrong...")
	} 

	// book slice to hold all books
	var bookArr[]models.Book

	defer bookList.Close()

	// look thru all the books in bookList and append to bookArr
	for bookList.Next() {
		var b models.Book
		if err := bookList.Scan(&b.ID, &b.Title, &b.Author, &b.Quantity); err != nil {
				c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Something went wrong..."})
		}
		bookArr = append(bookArr, b)
	}

	c.IndentedJSON(http.StatusOK, bookArr)
}

func GetBook(c *gin.Context) {
	// get id parameters
	id := c.Param("id")

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found..."})
	}

	c.IndentedJSON(http.StatusOK, book)
}

func PostBook(c *gin.Context) {
	body := models.Book{}
	data, err := c.GetRawData()

	if err != nil {
		c.AbortWithStatusJSON(400, "Something went wrong...")
		return
	}

	err = json.Unmarshal(data, &body)
	if err != nil {
		c.AbortWithStatusJSON(400, "Bad Input")
		return
	}

	// Use Exec for insert, update and delete queries
	_, err = database.Db.Exec("INSERT INTO books VALUES(DEFAULT, $1, $2, $3)", body.Title, body.Author, body.Quantity)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(400, "Could not create new Book")
	}else{
		c.IndentedJSON(http.StatusOK, body)
	}
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found..."})
	}
	// Use Exec for insert, update and delete queries
	_, err = database.Db.Exec("DELETE FROM books WHERE id=$1", id)

	if err != nil {
		fmt.Println("Error: ", err)
		c.AbortWithStatusJSON(400, "Could not delete book")
	}else{
		c.IndentedJSON(http.StatusOK, book)
	}

	// for index, book := range books {
	// 	if book.ID == id {
	// 		copy(books[index:], books[index+1:]) // shift all right elements of the deleted element to left one space
	// 		books = books[:len(books)-1] // truncate slice

	// 		c.IndentedJSON(http.StatusOK, book)
	// 		return
	// 	}
	// }

	// // return error message if book is not found
	// c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found..."})
}

func UpdateBook(c *gin.Context) {

	body := models.Book{}
	data, err := c.GetRawData()
	id := c.Param("id")

	if err != nil {
		c.AbortWithStatusJSON(400, "Something went wrong...")
		return
	}

	err = json.Unmarshal(data, &body)
	if err != nil {
		c.AbortWithStatusJSON(400, "Bad Input")
		return
	}

	// Use Exec for insert, update and delete queries
	_, err = database.Db.Exec("UPDATE books SET title=$1, author=$2, quantity=$3 WHERE id=$4", body.Title, body.Author, body.Quantity, id)

	if err != nil {
		fmt.Println("Error: ", err)
		c.AbortWithStatusJSON(400, "Could not update Book")
	}else{
		book, _ := getBookById(id)
		fmt.Println("here...")
		c.IndentedJSON(http.StatusOK, book)
	}

	// id := c.Param("id")
	// var newBook book

	// // BindJson to bind the receive JSON from body to newBook
	// if err := c.BindJSON(&newBook) ; err != nil {
	// 	return
	// }

	// bookData, err := getBookById(id)

	// if err != nil {
	// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
	// 	return
	// }

	// bookData.Title = newBook.Title
	// bookData.Author = newBook.Author
	// bookData.Quantity = newBook.Quantity

	// c.IndentedJSON(http.StatusOK, bookData)
}