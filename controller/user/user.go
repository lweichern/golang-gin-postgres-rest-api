package user

import (
	"example/http-server/controller"
	"example/http-server/database"
	"example/http-server/lib"
	"example/http-server/models"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

// USER CONTROLLER
func RegisterUser(c *gin.Context){
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid Request body: %s", err.Error())})
		return
	}

	// Hash password before store in DB
	hashedPassword := lib.HashPassword(user.Password)
	user.Password = hashedPassword

	database.Db.Create(&user)

	c.IndentedJSON(http.StatusOK, gin.H{"message": "User successfully created!", "body": user})
}

func LoginUser(c *gin.Context) {
	var loginUser models.User

	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid Request body: %s", err.Error())})
		return
	}

	// Find user in database by username
	var user models.User
	result := database.Db.Where("username=?", loginUser.Username).First(&user)

	// Hash password
	loginUser.Password = lib.HashPassword(loginUser.Password)

	// Chck if user exists in database, if exists, check password
	if result.Error != nil || user.Password != loginUser.Password {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Generate a unique token for the user
	token := uuid.New().String()

	fmt.Println("user: ", user)
	// Store user with newly generated token in users array
	controller.Users[token] = user

	// Return authenticated token to the client
	c.IndentedJSON(http.StatusOK, gin.H{"token": token})
}