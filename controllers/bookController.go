package controllers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Book struct {
	BookID string `json:"bookId"`
	Title  string `json:title`
	Author string	`json:author`
	Desc   string	`json:desc`
}

var BooksData = []Book{}

func CreateBook(ctx *gin.Context) {
	var newBook Book 
	// err := ctx.ShouldBindJSON(&newBook)
	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook.BookID = fmt.Sprintf("b%d", len(BooksData) + 1)
	BooksData = append(BooksData, newBook)

	ctx.JSON(http.StatusCreated, gin.H{
		"success": "true",
		"message": "Successfully add a book!",
		"data": newBook,
	})
}

func GetBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": BooksData,
		"message": "Successfully retrive all data!",
	})
}

func GetDetailBook(ctx *gin.Context) {
	bookId := ctx.Param("bookId")
	condition := false
	var result Book 

	for i, v := range BooksData {
		if bookId == v.BookID {
			condition = true
			result = BooksData[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"success": "false",
			"message": fmt.Sprintf("Book with id %v not found!", bookId),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": "true",
		"data": result,
		"message": "Successfully retrive detail a book!",
	})
}

func UpdateBook(ctx *gin.Context) {
	bookId := ctx.Param("bookId")
	condition := false
	var updatedBook Book 

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, v := range BooksData {
		if bookId == v.BookID {
			condition = true
			BooksData[i] = updatedBook
			BooksData[i].BookID = bookId
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"success": "false",
			"message": fmt.Sprintf("Book with id %v not found!", bookId),
		})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": fmt.Sprintf("Successfully update a book with id %v", bookId),
	})
}

func DeleteBook(ctx *gin.Context) {
	bookId := ctx.Param("bookId")
	condition := false
	var deleteBook int
	
	for i, book := range BooksData {
		if bookId == book.BookID {
			condition = true
			deleteBook = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"success": "false",
			"message": fmt.Sprintf("Book with id %v not found!", bookId),
		})
		return
	}

	copy(BooksData[deleteBook:], BooksData[deleteBook+1:])
	BooksData[len(BooksData)-1] = Book{}
	BooksData = BooksData[:len(BooksData)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"success": "true", 
		"message": "Successfully delete data!",
	})
}