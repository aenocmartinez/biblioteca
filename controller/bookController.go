package controller

import (
	"biblioteca/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RequestCreate struct {
	Title           string `json:"title" binding:"required"`
	Author          string `json:"author" binding:"required"`
	YearPublication int    `json:"year_publication" binding:"required"`
	Summary         string `json:"summary" binding:"required"`
}

type RequestUpdate struct {
	Id              int64  `json:"id" binding:"required"`
	Title           string `json:"title" binding:"required"`
	Author          string `json:"author" binding:"required"`
	YearPublication int    `json:"year_publication" binding:"required"`
	Summary         string `json:"summary" binding:"required"`
}

func BookList(c *gin.Context) {
	books := model.BookList()
	c.JSON(http.StatusOK, books)
}

func CreateBook(c *gin.Context) {
	var req RequestCreate
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	book := model.Book{}
	book.Title = req.Title
	book.Author = req.Author
	book.YearPublication = req.YearPublication
	book.Summary = req.Summary

	err = book.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "resource created successfully"})

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
	var req RequestUpdate
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

	c.JSON(http.StatusOK, gin.H{"message": "resource updated successfully"})
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

	c.JSON(http.StatusOK, gin.H{"message": "resource deleted successfully"})
}
