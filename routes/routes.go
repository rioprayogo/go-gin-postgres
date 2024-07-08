package routes

import (
	"go-gin-postgres"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Endpoint untuk buku
	r.GET("/books", controllers.GetAllBooks)
	// r.POST("/books", controllers.CreateBook)
	// r.GET("/books/:id", controllers.GetBookByID)
	// r.PUT("/books/:id", controllers.UpdateBook)
	// r.DELETE("/books/:id", controllers.DeleteBook)

	return r
}
