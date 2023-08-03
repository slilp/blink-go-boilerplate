package models

import (
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
	Products 	[]ProductEntity `gorm:"many2many:order_product;" json:"products"` 
}


func (OrderEntity) TableName() string {
    return "orders"
}

