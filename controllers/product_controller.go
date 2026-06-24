package controllers

import (
	"net/http"

	"github.com/andy-kajeke/goDemoApi/config"
	"github.com/andy-kajeke/goDemoApi/models"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"info": gin.H{
				"code":    400,
				"message": err.Error(),
			},
		})
		return
	}

	if err := config.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"info": gin.H{
				"code":    500,
				"message": err.Error(),
			},
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"info": gin.H{
			"code": 201,
			"data": product,
		},
	})
}

func GetProducts(c *gin.Context) {
	var products []models.Product

	config.DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"info": gin.H{
			"code": 200,
			"data": products,
		},
	})
}
func GetProductByID(c *gin.Context) {
	id := c.Param("id")

	var product models.Product

	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "failed",
			"info": gin.H{
				"code":    404,
				"message": "Product not found invalid Id",
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"info": gin.H{
			"code": 200,
			"data": product,
		},
	})
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	var product models.Product

	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "failed",
			"info": gin.H{
				"code":    404,
				"message": "Product not found invalid Id",
			},
		})
		return
	}

	var input models.Product

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"info": gin.H{
				"code":    400,
				"message": err.Error(),
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

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"info": gin.H{
			"code": 201,
			"data": product,
		},
	})
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	var product models.Product

	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "failed",
			"info": gin.H{
				"code":    404,
				"message": "Product not found invalid Id",
			},
		})
		return
	}

	config.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"info": gin.H{
			"code":    200,
			"message": "Product deleted successfully",
		},
	})
}
