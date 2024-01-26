package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gopass/models"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")
	router.Static("/static", "./static")

	router.GET("/", func(c *gin.Context) { 
		c.HTML(http.StatusOK,
			"index.html",
			gin.H{"title": "Home Page"},
			)
  	})
	router.GET("/home", func(c *gin.Context) { 
		c.HTML(http.StatusOK,
			"home.html",
			gin.H{"title": "Home Page"},
			)
  	})
	router.GET("/items", models.GetAllItems)

	models.ConnectDB()

	router.Run()
}