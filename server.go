package main

import (
	"example/http-server/database"
	"example/http-server/routes"

	"github.com/gin-gonic/gin"
)


func main(){
	router := gin.Default()
	database.ConnectDatabase()
	routes.BookRoute(router)
	router.Run("localhost:8080")
}