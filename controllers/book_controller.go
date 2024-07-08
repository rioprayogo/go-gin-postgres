package controllers

import (
	"net/http"

	"your-module/repositories"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(c *gin.Context) {
	books, err := repositories.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}

	c.JSON(http.StatusOK, books)
}

// Implementasi lainnya seperti CreateBook, GetBookByID, UpdateBook, dan DeleteBook
