package migration

import (
	order "blink-go-gin-boilerplate/app/order/api"
	product "blink-go-gin-boilerplate/app/product/api"
	user "blink-go-gin-boilerplate/app/user/api"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&product.ProductEntity{})
	db.AutoMigrate(&user.UserEntity{})
	db.AutoMigrate(&order.OrderEntity{})
}