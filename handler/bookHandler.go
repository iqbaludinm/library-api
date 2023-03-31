package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iqbaludinm/library-api/helpers"
	"github.com/iqbaludinm/library-api/models"
)

func (h HttpServer) CreateEmployee(c *gin.Context) {
	req := models.Book{}
	err := c.BindJSON(&req)
	if err != nil {
		helpers.BadRequest(c, err.Error())
		return
	}

	v := req.Validate()
	if v != nil {
		helpers.ErrorWithData(c, v.Error())
		return
	}

	// call service
	res, err := h.app.CreateBook(req)
	if err != nil {
		helpers.InternalServerError(c, err.Error())
		return
	}
	// helpers.OkWithData(c, "Success Add Data!", res)
	helpers.Ok(c, res)
}

func (h HttpServer) GetBooks(c *gin.Context) {
	res, err := h.app.GetBooks()
	if err != nil {
		helpers.InternalServerError(c, err.Error())
		return
	}

	if res == nil {
		helpers.OkWithData(c, "Success Retrive All Data", []models.Book{})
	} else {
		helpers.OkWithData(c, "Success Retrive All Data", res)
	}
}

func (h HttpServer) GetBookById(c *gin.Context) {
	req, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.BadRequest(c, "Bad Parameter", err)
		return
	}

	res, err := h.app.GetBookById(int64(req))
	if err != nil {
		helpers.NotFound(c, err.Error())
		return
	}
	
	// helpers.OkWithData(c, "Success Retrive a Data", res)
	helpers.Ok(c, res)
}

func (h HttpServer) UpdateBook(c *gin.Context) {
	req := models.Book{}
	err := c.BindJSON(&req)
	if err != nil {
		helpers.BadRequest(c, err.Error())
		return
	}

	param, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.BadRequest(c, "Bad Parameter", err)
		return
	}

	_, e := h.app.GetBookById(int64(param))
	if e != nil {
		helpers.NotFound(c, e.Error())
		return
	}

	req.ID = param

	res, err := h.app.UpdateBook(req)
	if err != nil {
		helpers.ErrorWithData(c, err)
		return
	}
	helpers.Ok(c, res)
}

func (h HttpServer) DeleteBook(c *gin.Context) {
	req, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.BadRequest(c, "Bad Parameter", err)
		return
	}


	_, er := h.app.DeleteBook(int64(req))
	if er != nil {
		
		helpers.NotFound(c, er.Error())
		return
	}
	helpers.OkWithMessage(c, "Book deleted successfully")
}
