package routes

import (
	"example/http-server/controller"

	"github.com/gin-gonic/gin"
)

func BookRoute(router *gin.Engine){
	router.GET("/books", controller.GetBooks)
	router.GET("/books/:id", controller.GetBook)
	router.POST("/books", controller.PostBook)
	router.DELETE("/books/:id", controller.DeleteBook)
	router.PATCH("/books/:id", controller.UpdateBook)
}