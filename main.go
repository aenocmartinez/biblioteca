package main

import (
	"biblioteca/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())
	r.GET("books", controller.BookList)
	r.GET("books/:id", controller.ReadBook)
	r.POST("books", controller.CreateBook)
	r.DELETE("books/:id", controller.DeleteBook)
	r.PUT("books", controller.UpdateBook)

	r.Run(":8080")
}
