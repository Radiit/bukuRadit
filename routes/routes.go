package routes

import (
	"bukuRadit/controllers"

	"github.com/gin-gonic/gin"
)

func BookRoute(router *gin.Engine) {
	router.GET("/books", controllers.GetAllBooks)
	router.GET("/books/:id", controllers.GetBookByID)
	router.POST("/books", controllers.CreateBook)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)
}
