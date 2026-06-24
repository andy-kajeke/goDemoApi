package migrations

import (
	"github.com/andy-kajeke/goDemoApi/config"
	"github.com/andy-kajeke/goDemoApi/models"
)

func MigrateModels() {
	err := config.DB.AutoMigrate(
		&models.User{},
		&models.Product{},
	)
	if err != nil {
		panic("Failed to run database migrations: " + err.Error())
	}
}
