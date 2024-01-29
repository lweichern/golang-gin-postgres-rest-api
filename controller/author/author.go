package author

import (
	"example/http-server/database"
	"example/http-server/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var DB = database.Db

func GetAuthors(c *gin.Context){
	var authors []models.Author

	// Preload to load all the books associated to the authors
	result := database.Db.Preload("Books").Find(&authors)

	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": result.Error})
		return
	}

	c.IndentedJSON(http.StatusOK, authors)
}

func GetAuthor(c *gin.Context) {
	id := c.Param("id")
	var author models.Author

	// Preload to load all the books associated to the authors
	result := database.Db.Preload("Books").First(&author, id)

	if result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": result.Error})
		return
	}

	c.IndentedJSON(http.StatusOK, author)
}

func PostAuthor(c *gin.Context) {
	var author models.Author

	if err := c.ShouldBindJSON(&author); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid request body: %s", err.Error())})
		return
	}

	database.Db.Create(&author)

	c.IndentedJSON(http.StatusOK, gin.H{"messaage": "Author created successfully!", "author": author})
}