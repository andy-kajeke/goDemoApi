package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" gorm:"unique" binding:"required,email"`
	Phone    string `json:"phone"`
	Username string `json:"username" gorm:"unique" binding:"required"`
}
