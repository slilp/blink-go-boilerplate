package models

import (
	"encoding/json"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RoleType string

const (
    USER  RoleType = "USER"
    ADMIN RoleType = "ADMIN"
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
	Role		RoleType `sql:"type:ENUM('USER', 'ADMIN')" gorm:"default:'USER'" json:"role"`
	IsActive	bool   `gorm:"not null; default:true" json:"isActive"`
	Orders 		[]OrderEntity `gorm:"foreignKey:UserID" json:"orders"`
}

func (u *UserEntity) GenerateEncryptedPassword() string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	return string(hash)
}

func (u *UserEntity) ValidateEncryptedPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
    return err == nil
}

func  (u *UserEntity) GenerateJwtToken() (string,string) {

	user,_ := json.Marshal(u)
	payload := string(user)
	accessToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(time.Now().Add(10 * time.Minute).Unix(), 0)),
		Subject:  payload,
	}).SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	refreshToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Unix(time.Now().Add(24 * time.Hour).Unix(), 0)),
	Subject:  payload,
	}).SignedString([]byte(os.Getenv("REFRESH_JWT_SECRET_KEY")))
	
		
	return  accessToken	,refreshToken
}

func (UserEntity) TableName() string {
    return "users"
}

