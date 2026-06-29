package controllers

import (
	"net/http"
	"strings"

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
	page := c.GetInt("page")
	limit := c.GetInt("limit")
	offset := c.GetInt("offset")
	search := strings.TrimSpace(c.GetString("search"))

	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	var users []models.User
	var total int64
	query := config.DB.Model(&models.User{})

	if search != "" {
		searchPattern := "%" + search + "%"
		query = query.Where("name ILIKE ? OR email ILIKE ? OR username ILIKE ?", searchPattern, searchPattern, searchPattern)
	}

	startDate, hasStartDate := c.Get("startDate")
	endDate, hasEndDate := c.Get("endDate")

	if hasStartDate && hasEndDate {
		query = query.Where("created_at BETWEEN ? AND ?", startDate, endDate)
	} else if hasStartDate {
		query = query.Where("created_at >= ?", startDate)
	} else if hasEndDate {
		query = query.Where("created_at <= ?", endDate)
	}

	if err := query.Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, middleware.APIResponse{
			Status: "Failed",
			Info: middleware.ResponseInfo{
				Code:    500,
				Message: err.Error(),
			},
		})
		return
	}

	if err := query.
		Select("id", "created_at", "updated_at", "name", "email", "phone", "username").
		Order("created_at DESC, id DESC").
		Limit(limit).
		Offset(offset).
		Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, middleware.APIResponse{
			Status: "Failed",
			Info: middleware.ResponseInfo{
				Code:    500,
				Message: err.Error(),
			},
		})
		return
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))

	c.JSON(http.StatusOK, middleware.APIResponse{
		Status: "Success",
		Info: middleware.ResponseInfo{
			Code:    200,
			Message: "Users fetched successfully",
			Pagination: middleware.ResponsePagination{
				Page:        page,
				Limit:       limit,
				Total:       total,
				TotalPages:  totalPages,
				HasNextPage: total > int64(page*limit),
			},
			Data: users,
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
