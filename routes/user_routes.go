package routes

import (
	"github.com/andy-kajeke/goDemoApi/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	userRoutes := r.Group("/api/users")
	{
		userRoutes.POST("/", controllers.CreateUser)
		userRoutes.GET("/", controllers.GetUsers)
		userRoutes.GET("/:id", controllers.GetUserByID)
		userRoutes.PUT("/:id", controllers.UpdateUser)
		userRoutes.DELETE("/:id", controllers.DeleteUser)
	}
}
