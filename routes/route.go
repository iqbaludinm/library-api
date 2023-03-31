package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iqbaludinm/library-api/handler"
)

func RegisterAPI(r *gin.Engine, server handler.HttpServer) {
	apiBook := r.Group("/api/v1/books")
	{
		apiBook.POST("", server.CreateEmployee)
		apiBook.GET("", server.GetBooks)
		apiBook.GET("/:id", server.GetBookById)
		apiBook.PUT("/:id", server.UpdateBook)
		apiBook.DELETE("/:id", server.DeleteBook)
	}

}
