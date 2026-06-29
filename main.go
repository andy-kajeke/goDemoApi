package main

import (
	"net/http"
	"os"

	"github.com/andy-kajeke/goDemoApi/config"
	"github.com/andy-kajeke/goDemoApi/internal"
	"github.com/andy-kajeke/goDemoApi/middleware"
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
		c.JSON(http.StatusOK, middleware.APIResponse{
			Status: "Success",
			Info: middleware.ResponseInfo{
				Code:    200,
				Message: "Gin demo CRUD API is running",
			},
		})
	})

	r.Run(":" + os.Getenv("APP_PORT"))
}
