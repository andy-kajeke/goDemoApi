package main

import (
	"github.com/andy-kajeke/goDemoApi/config"
	"github.com/andy-kajeke/goDemoApi/models"
	"github.com/andy-kajeke/goDemoApi/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDatabase()

	config.DB.AutoMigrate(&models.User{})

	routes.UserRoutes(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Gin CRUD API is running",
		})
	})

	r.Run(":8080")
}
