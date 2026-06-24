package main

import (
	"os"

	"github.com/andy-kajeke/goDemoApi/config"
	"github.com/andy-kajeke/goDemoApi/internal"
	"github.com/andy-kajeke/goDemoApi/migrations"
	"github.com/andy-kajeke/goDemoApi/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	internal.LoadEnv()

	config.ConnectDatabase()

	migrations.MigrateModels()

	//config.DB.AutoMigrate(&models.User{})

	routes.SystemRoutes(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Gin demo CRUD API is running",
		})
	})

	r.Run(":" + os.Getenv("APP_PORT"))
}
