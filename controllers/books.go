package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/simphiwehlabisa/go-crud-api/models"
)

// FindBook controller
func FindBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//get model if exists
	var book models.Book
	if err := db.Where("id = ? ", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Nod Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// CreateBookInput Schema
type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// UpdateBookInput Schema
type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// UpdateBook
func UpdateBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//check if model exists
	var book models.Book
	if err := db.Where("id = ? ", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	//validate input
	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&book).Update(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
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

// DeleteBook controller
func DeleteBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//check model if its valid
	var book models.Book

	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	db.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
