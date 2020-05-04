package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/simphiwehlabisa/go-crud-api/controllers"
	"github.com/simphiwehlabisa/go-crud-api/models"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	router := gin.Default()
	db := models.SetupModels()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"title": "Home Page",
			},
		)
	})

	//api routes
	api := router.Group("/api")
	api.GET("/books", controllers.FindBooks)
	api.POST("/books", controllers.CreateBook)
	api.GET("/books/:id", controllers.FindBook)
	api.PATCH("/books/:id", controllers.UpdateBook)
	api.DELETE("/books/:id", controllers.DeleteBook)

	dotenv := goDotEnvVariable("APP_ENV")

	if dotenv == "local" {
		router.Run("localhost:8082")
	} else {
		router.Run()

	}

	// router.Run()

}
