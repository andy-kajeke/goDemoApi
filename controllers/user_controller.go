package controllers

import (
	"net/http"

	"github.com/andy-kajeke/goDemoApi/config"
	"github.com/andy-kajeke/goDemoApi/middleware"
	"github.com/andy-kajeke/goDemoApi/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, middleware.APIResponse{
			Status: "Failed",
			Info: middleware.ResponseInfo{
				Code:    400,
				Message: err.Error(),
			},
		})
		return
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, middleware.APIResponse{
			Status: "Failed",
			Info: middleware.ResponseInfo{
				Code:    500,
				Message: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusCreated, middleware.APIResponse{
		Status: "Success",
		Info: middleware.ResponseInfo{
			Code: 201,
			Data: user,
		},
	})
}

func GetUsers(c *gin.Context) {
	var users []models.User

	config.DB.Find(&users)

	c.JSON(http.StatusOK, middleware.APIResponse{
		Status: "Success",
		Info: middleware.ResponseInfo{
			Code:    200,
			Message: "Users fetched successfully",
			Data:    users,
		},
	})
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, middleware.APIResponse{
			Status: "Failed",
			Info: middleware.ResponseInfo{
				Code:    404,
				Message: "User not found invalid Id",
			},
		})
		return
	}

	c.JSON(http.StatusOK, middleware.APIResponse{
		Status: "Success",
		Info: middleware.ResponseInfo{
			Code:    200,
			Message: "User fetched successfully",
			Data:    user,
		},
	})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, middleware.APIResponse{
			Status: "Failed",
			Info: middleware.ResponseInfo{
				Code:    404,
				Message: "User not found invalid Id",
			},
		})
		return
	}

	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, middleware.APIResponse{
			Status: "Failed",
			Info: middleware.ResponseInfo{
				Code:    400,
				Message: err.Error(),
			},
		})
		return
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Phone = input.Phone
	user.Username = input.Username

	config.DB.Save(&user)

	c.JSON(http.StatusOK, middleware.APIResponse{
		Status: "Success",
		Info: middleware.ResponseInfo{
			Code:    200,
			Message: "User record updated successfully",
			Data:    user,
		},
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, middleware.APIResponse{
			Status: "Failed",
			Info: middleware.ResponseInfo{
				Code:    404,
				Message: "User not found invalid Id",
			},
		})
		return
	}

	config.DB.Delete(&user)

	c.JSON(http.StatusOK, middleware.APIResponse{
		Status: "Success",
		Info: middleware.ResponseInfo{
			Code:    404,
			Message: "User deleted successfully",
		},
	})
}
