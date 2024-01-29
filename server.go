package main

import (
	"example/http-server/database"
	"example/http-server/middleware"
	"example/http-server/routes"

	"github.com/gin-gonic/gin"
)


func main(){
	router := gin.Default()
	database.ConnectDatabase()
	routes.UserRoute(router)
	router.Use(middleware.AuthMiddleware) // any route below this will be protected routes
	routes.AuthorRoute(router)
	routes.BookRoute(router)
	router.Run("localhost:8080")
}