package routes

import (
	"example/http-server/controller/author"
	"example/http-server/controller/book"
	"example/http-server/controller/user"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine){
	router.POST("/register", user.RegisterUser)
	router.POST("/login", user.LoginUser)
}

func AuthorRoute(router *gin.Engine){
	authorRoutes := router.Group("/authors")
	{
		authorRoutes.GET("/", author.GetAuthors)
		authorRoutes.GET("/:id", author.GetAuthor)
		authorRoutes.POST("/", author.PostAuthor)
	}
}

func BookRoute(router *gin.Engine){
	bookRoute := router.Group("/books")
	{
		bookRoute.GET("/", book.GetBooks)
		bookRoute.GET("/:id", book.GetBook)
		bookRoute.POST("/", book.PostBook)
		bookRoute.DELETE("/:id", book.DeleteBook)
		bookRoute.PATCH("/:id", book.UpdateBook)
	}
}