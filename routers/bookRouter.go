package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/iqbaludinm/library-api/controllers"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.POST("/books", controllers.CreateBook)
	router.GET("/books", controllers.GetBooks)
	router.GET("/books/:bookId", controllers.GetDetailBook)
	router.PUT("/books/:bookId", controllers.UpdateBook)
	router.DELETE("/books/:bookId", controllers.DeleteBook)

	return router
}