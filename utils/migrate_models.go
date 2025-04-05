package utils

import (
	"tripat3k2/url_shortner/models"

	"gorm.io/gorm"
)

func MigrateModels(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Url{})
}
