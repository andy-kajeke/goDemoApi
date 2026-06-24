package controllers

import (
	"net/http"

	"github.com/andy-kajeke/goDemoApi/config"
	"github.com/andy-kajeke/goDemoApi/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   user,
	})
}

func GetUsers(c *gin.Context) {
	var users []models.User

	config.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"info": gin.H{
			"code": 200,
			"data": users,
		},
	})
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "failed",
			"message": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "failed",
			"message": "User not found",
		})
		return
	}

	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Phone = input.Phone
	user.Username = input.Username

	config.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user,
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "failed",
			"message": "User not found",
		})
		return
	}

	config.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "User deleted successfully",
	})
}
