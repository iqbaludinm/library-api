package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	// SUCCESS_GET_ALL_DATA = "Successfully Retrieve All Data"
	// SUCCESS_GET_DATA = "Successfully Retrieve A Data"
	ERR_NOT_FOUND = "Data not found!"
)

func OkWithMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK, 
		"message": message,
	})
}

func OkWithData(c *gin.Context, message interface{}, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK, 
		"message": message,
		"data": data,
	})
}

func NoContent(c *gin.Context) {
	c.JSON(http.StatusCreated, nil)
}

func BadRequest(c *gin.Context, message string, data ...interface{}) {
	obj := gin.H{"status": http.StatusBadRequest, "message": message}
	if len(data) > 0 {
		obj["data"] = data[0]
	}
	c.JSON(http.StatusBadRequest, obj)
}

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": ERR_NOT_FOUND})
}

func InternalServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": message})
}

func ErrorWithData(c *gin.Context, err interface{}) {
	c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err})
}