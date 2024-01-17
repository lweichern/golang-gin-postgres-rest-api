package controller

import (
	// "encoding/json"
	// "errors"
	"example/http-server/database"
	"example/http-server/models"
	"fmt"
	"net/http"

	// "crypto/rand"
	// "encoding/hex"

	"github.com/gin-gonic/gin"
)

// tokens slice to store token generated for verfication purpose
var tokens []string

func GetBooks(c *gin.Context){
	var bookList []models.Book
	// Find books
	result := database.Db.Find(&bookList)

	// Check errors
	if result.Error != nil {
		fmt.Println("Error finding books: ", result.Error)
		c.AbortWithStatusJSON(400, "Something went wrong...")
	} 

	c.IndentedJSON(http.StatusOK, bookList)
}

func GetBook(c *gin.Context) {
	// get id parameters
	id := c.Param("id")

	var book models.Book

	// Find book based on ID
	result := database.Db.First(&book, id)

	// Check if book exists
	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found..."})
	}

	c.IndentedJSON(http.StatusOK, book)
}

func PostBook(c *gin.Context) {
	var bookData models.Book

	// Bind the Json body to Book struct
	if err := c.ShouldBindJSON(&bookData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid request body: %s", err.Error())})
		return
	}

	// Create book in database
	database.Db.Create(&bookData)

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Book created successfully", "book": bookData})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book

	// Find book
	result := database.Db.First(&book, id)

	// Check if book exists
	if result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not found..."})
		return
	}

	// Delete book from db
	database.Db.Delete(&book)

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Book deleted successfully", "book": book})
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book

	// Find book based on ID
	result := database.Db.First(&book, id)

	// Check if book exists
	if result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not found..."})
		return
	}

	// Bind JSON body to the User Struct for partial update
	if err := c.ShouldBindJSON(&book); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid request body: %s", err.Error())})
		return
	}

	// Update the book in the database
	database.Db.Save(&book)

	// Return updated book to the client
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Book updated successfully!", "book": book})
}