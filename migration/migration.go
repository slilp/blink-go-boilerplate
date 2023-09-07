package migration

import (
	order "github.com/slilp/blink-go-boilerplate/app/order/api"
	product "github.com/slilp/blink-go-boilerplate/app/product/api"
	user "github.com/slilp/blink-go-boilerplate/app/user/api"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&product.ProductEntity{})
	db.AutoMigrate(&user.UserEntity{})
	db.AutoMigrate(&order.OrderEntity{})
}