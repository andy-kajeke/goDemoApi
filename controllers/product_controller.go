package controllers

import (
	"net/http"
	"strings"

	"github.com/andy-kajeke/goDemoApi/config"
	"github.com/andy-kajeke/goDemoApi/middleware"
	"github.com/andy-kajeke/goDemoApi/models"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, middleware.APIResponse{
			Status: "Failed",
			Info: middleware.ResponseInfo{
				Code:    400,
				Message: err.Error(),
			},
		})
		return
	}

	if err := config.DB.Create(&product).Error; err != nil {
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
			Data: product,
		},
	})
}

func GetProducts(c *gin.Context) {
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

	var products []models.Product
	var total int64
	query := config.DB.Model(&models.Product{})

	if search != "" {
		searchPattern := "%" + search + "%"
		query = query.Where("name ILIKE ? OR price ILIKE ?", searchPattern, searchPattern)
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

	if err := query.
		Select("id", "created_at", "updated_at", "name", "openingStock", "lowStockAlert", "price", "description").
		Order("created_at DESC, id DESC").
		Limit(limit).
		Offset(offset).
		Find(&products).Error; err != nil {
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

	//config.DB.Find(&products)

	c.JSON(http.StatusOK, middleware.APIResponse{
		Status: "Success",
		Info: middleware.ResponseInfo{
			Code:    200,
			Message: "Products fetched successfully",
			Pagination: middleware.ResponsePagination{
				Page:        page,
				Limit:       limit,
				Total:       total,
				TotalPages:  totalPages,
				HasNextPage: total > int64(page*limit),
			},
			Data: products,
		},
	})
}
func GetProductByID(c *gin.Context) {
	id := c.Param("id")

	var product models.Product

	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, middleware.APIResponse{
			Status: "Failed",
			Info: middleware.ResponseInfo{
				Code:    404,
				Message: "Product not found invalid Id",
			},
		})
		return
	}

	c.JSON(http.StatusOK, middleware.APIResponse{
		Status: "Success",
		Info: middleware.ResponseInfo{
			Code:    200,
			Message: "Product fetched successfully",
			Data:    product,
		},
	})
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	var product models.Product

	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, middleware.APIResponse{
			Status: "Failed",
			Info: middleware.ResponseInfo{
				Code:    404,
				Message: "Product not found invalid Id",
			},
		})
		return
	}

	var input models.Product

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

	product.Name = input.Name
	product.OpeningStock = input.OpeningStock
	product.LowStockAlert = input.LowStockAlert
	product.Price = input.Price
	product.Description = input.Description

	config.DB.Save(&product)

	c.JSON(http.StatusOK, middleware.APIResponse{
		Status: "Success",
		Info: middleware.ResponseInfo{
			Code:    200,
			Message: "Product record updated successfully",
			Data:    product,
		},
	})
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	var product models.Product

	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, middleware.APIResponse{
			Status: "Failed",
			Info: middleware.ResponseInfo{
				Code:    404,
				Message: "Product not found invalid Id",
			},
		})
		return
	}

	config.DB.Delete(&product)

	c.JSON(http.StatusOK, middleware.APIResponse{
		Status: "Success",
		Info: middleware.ResponseInfo{
			Code:    404,
			Message: "Product deleted successfully",
		},
	})
}
