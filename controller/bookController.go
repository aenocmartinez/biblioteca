package controller

import (
	"biblioteca/model"
	"biblioteca/view"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BookList(c *gin.Context) {
	books := model.BookList()
	c.JSON(http.StatusOK, books)
}

func CreateBook(c *gin.Context) {
	var req view.RequestCreate
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	book := model.NewBook(req.Title, req.Author)
	book.YearPublication = req.YearPublication
	book.Summary = req.Summary

	err = book.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, book)

}

func ReadBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := model.ReadBook(int64(id))
	if !book.Exists() {
		c.JSON(http.StatusNotFound, gin.H{"message": "resource not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	var req view.RequestUpdate
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	book := model.ReadBook(req.Id)
	if !book.Exists() {
		c.JSON(http.StatusNotFound, gin.H{"message": "resource not found"})
		return
	}

	book.Title = req.Title
	book.Author = req.Author
	book.YearPublication = req.YearPublication
	book.Summary = req.Summary

	err = book.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := model.ReadBook(int64(id))
	if !book.Exists() {
		c.JSON(http.StatusNotFound, gin.H{"message": "resource not found"})
		return
	}

	err = book.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted resource"})
}
