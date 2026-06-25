package controllers

import (
	"net/http"

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
	var products []models.Product

	config.DB.Find(&products)

	c.JSON(http.StatusOK, middleware.APIResponse{
		Status: "Success",
		Info: middleware.ResponseInfo{
			Code:    200,
			Message: "Products fetched successfully",
			Data:    products,
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
