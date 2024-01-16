package routes

import (
	"example/http-server/controller"

	"github.com/gin-gonic/gin"
)

func BookRoute(router *gin.Engine){
	router.GET("/books", controller.GetBooks)
	router.POST("/books", controller.PostBook)
	router.GET("/books/:id", controller.GetBook)
	router.DELETE("/books/:id", controller.DeleteBook)
	router.PUT("/books/:id", controller.UpdateBook)
}