package routes

import (
	"github.com/andy-kajeke/goDemoApi/controllers"
	"github.com/gin-gonic/gin"
)

func SystemRoutes(r *gin.Engine) {
	base_url := "/api"

	userRoutes := r.Group(base_url + "/users")
	{
		userRoutes.POST("/", controllers.CreateUser)
		userRoutes.GET("/", controllers.GetUsers)
		userRoutes.GET("/:id", controllers.GetUserByID)
		userRoutes.PUT("/:id", controllers.UpdateUser)
		userRoutes.DELETE("/:id", controllers.DeleteUser)
	}

	productRoutes := r.Group(base_url + "/products")
	{
		productRoutes.POST("/", controllers.CreateProduct)
		productRoutes.GET("/", controllers.GetProducts)
		productRoutes.GET("/:id", controllers.GetProductByID)
		productRoutes.PUT("/:id", controllers.UpdateProduct)
		productRoutes.DELETE("/:id", controllers.DeleteProduct)
	}
}
