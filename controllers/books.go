package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/simphiwehlabisa/go-crud-api/models"
)

// CreateBookInput Schema
type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// CreateBook controller
func CreateBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//validate input
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//create book
	book := models.Book{Title: input.Title, Author: input.Author}
	db.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// FindBooks controller
// get all books
func FindBooks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var books []models.Book
	db.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}
