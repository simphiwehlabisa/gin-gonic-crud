package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/simphiwehlabisa/go-crud-api/models"
)

// GET /books
// get all books
func FindBooks(c *gin.Context) {
	db := c.Mustget("db").(*gorm.DB)

	var books []models.Book
	db.find(&books)

	c.Json(http.StatusOk, gin.H{"data": books})
}
