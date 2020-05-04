package main

import (
	"github.com/gin-gonic/gin"

	"github.com/simphiwehlabisa/go-crud-api/models"
)

func main() {
	router := gin.Default()
	db := models.SetupModels()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// router.LoadHTMLGlob("templates/*")
	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(
	// 		http.StatusOK,
	// 		"index.html",
	// 		gin.H{
	// 			"title": "Home Page",
	// 		},
	// 	)
	// })

	router.Run("localhost:8082")
}
