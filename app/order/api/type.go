package order

import (
	product "github.com/slilp/blink-go-boilerplate/app/product/api"

	"gorm.io/gorm"
)



type OrderStatus string

const (
	DRAFTED OrderStatus = "DRAFTED"
	PROCESSING  OrderStatus = "PROCESSING"
	DELIVERING OrderStatus = "DELIVERING"
	REJECTED OrderStatus = "REJECTED"
	COMPLETED OrderStatus = "COMPLETED"
)

type OrderEntity struct{
	gorm.Model
	Status		OrderStatus 	`sql:"type:ENUM('DRAFTED', 'PROCESSING', 'DELIVERING', 'REJECTED', 'COMPLETED')" gorm:"default:'DRAFTED'" json:"status"`
	UserID 		uint			`gorm:"column:user_id" json:"userId"`
	Products 	[]product.ProductEntity `gorm:"many2many:order_product;" json:"products"` 
}


func (OrderEntity) TableName() string {
    return "orders"
}


type CreateOrderRequest struct{
	Products  []product.ProductEntity `form:"products" json:"products" binding:"required"`
}


type UpdateOrderStatusRequest struct{
	Status   OrderStatus `form:"status" json:"status" binding:"required"`
}