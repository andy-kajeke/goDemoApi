package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name          string  `json:"name" binding:"required"`
	OpeningStock  int     `json:"openingStock" binding:"required"`
	LowStockAlert int     `json:"lowStockAlert" binding:"required"`
	Price         float32 `json:"price" binding:"required" gorm:"type:decimal(10,2)"`
	Description   string  `json:"description"`
}
