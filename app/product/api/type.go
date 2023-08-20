package product

import (
	"gorm.io/gorm"
)


type ProductEntity struct{
	gorm.Model
	Sku			string 	`gorm:"unique_index; not null;" json:"sku"`
	Name		string 	`gorm:"not null" json:"name"`
	Description	string 	`gorm:"column:description" json:"description"`
	Pic			string 	`gorm:"column:pic" json:"pic"`
	Price		float64 `gorm:"column:price default:0" json:"price"`
	IsAvailable	bool   	`gorm:"not null; default:true" json:"isAvailable"`
}


func (ProductEntity) TableName() string {
    return "products"
}


type CreateProductRequest struct {
	Sku   		string      `form:"sku" json:"sku" binding:"required"`
	Name    	string      `form:"name" json:"name" binding:"required"`
	Description	string 		`form:"description" json:"description"`
	Pic			string 		`form:"pic" json:"pic"`
	Price		float64 	`form:"price" json:"price" binding:"numeric,gte=0"`
}

type UpdateProductRequest struct {
	Sku   		string      `form:"sku" json:"sku"`
	Name    	string      `form:"name" json:"name"`
	Description	string 		`form:"description" json:"description"`
	Pic			string 		`form:"pic" json:"pic"`
	Price		float64 	`form:"price" json:"price" binding:"numeric,gte=0"`
}

