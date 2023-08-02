package migration

import (
	"blink-go-gin-boilerplate/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.UserEntity{})
	db.AutoMigrate(&models.ProductEntity{})
	db.AutoMigrate(&models.OrderEntity{})
	// db.AutoMigrate(&models.OrderProductEntity{})
}