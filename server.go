package main

import (
	"example/http-server/database"
	"example/http-server/middleware"
	"example/http-server/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	database.ConnectDatabase()
	routes.UserRoute(router)
	router.Use(middleware.AuthMiddleware) // any route below this will be protected routes
	routes.AuthorRoute(router)
	routes.BookRoute(router)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router.Run("0.0.0.0:" + port)
}
