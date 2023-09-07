package user

import (
	order "github.com/slilp/blink-go-boilerplate/app/order/api"
	"github.com/slilp/blink-go-boilerplate/utils"

	"gorm.io/gorm"
)

type UserEntity struct{
	gorm.Model
	Username	string `gorm:"unique_index; not null;" json:"username"`
	Password	string `gorm:"not null" json:"password"`
	FirstName	string `gorm:"column:first_name" json:"firstName"`
	LastName	string `gorm:"column:last_name" json:"lastName"`
	Email		string `gorm:"unique_index; not null" json:"email"`
	Avatar		string `gorm:"not null" json:"avatar"`
	Phone		string `gorm:"not null" json:"phone"`
	Role		utils.RoleType `sql:"type:ENUM('USER', 'ADMIN')" gorm:"default:'USER'" json:"role"`
	IsActive	bool   `gorm:"not null; default:true" json:"isActive"`
	Orders 		[]order.OrderEntity `gorm:"foreignKey:UserID" json:"orders"`
}

func (UserEntity) TableName() string {
    return "users"
}

type RegisterRequest struct {
	Username   	string      `form:"username" json:"username" binding:"required,email"`
	Password    string      `form:"password" json:"password" binding:"required,min=5"`
	FirstName	string 		`form:"firstName" json:"firstName"`
	LastName	string 		`form:"firstName" json:"lastName"`
	Email		string 		`form:"email" json:"email" binding:"required,email"`
	Avatar		string 		`form:"avatar" json:"avatar"`
	Phone		string 		`form:"phone" json:"phone"`
	Role		utils.RoleType `form:"role" json:"role" default:"USER"`
}

type SignInRequest struct {
	Username   	string      `form:"username" json:"username" binding:"required,email" example:"admin@email.com"`
	Password    string      `form:"password" json:"password" binding:"required,min=5" example:"admin1234"`
}