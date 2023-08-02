package controllers

import (
	"blink-go-gin-boilerplate/middleware"
	"blink-go-gin-boilerplate/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type RegisterRequest struct {
	Username   	string      `form:"username" json:"username" binding:"required,email"`
	Password    string      `form:"password" json:"password" binding:"required,min=5"`
	FirstName	string 		`form:"firstName" json:"firstName"`
	LastName	string 		`form:"firstName" json:"lastName"`
	Email		string 		`form:"email" json:"email" binding:"required,email"`
	Avatar		string 		`form:"avatar" json:"avatar"`
	Phone		string 		`form:"phone" json:"phone"`
	Role		models.RoleType `form:"role" json:"role" default:"USER"`
}

type SignInRequest struct {
	Username   	string      `form:"username" json:"username" binding:"required,email" example:"admin@email.com"`
	Password    string      `form:"password" json:"password" binding:"required,min=5" example:"admin1234"`
}

type User struct {
	DB *gorm.DB
}

// Register godoc
// @summary Register
// @tags user
// @id Register
// @accept json
// @produce json
// @param User body RegisterRequest true "''"
// @router /user/register [post]
func (u *User) Register(ctx *gin.Context) {
	var form RegisterRequest
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var checkUser models.UserEntity
	if check := u.DB.First(&checkUser,"username = ?", form.Username); check.RowsAffected != 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Duplicate"})
		return
	}


	var user models.UserEntity
	copier.Copy(&user, &form)

	user.Password = user.GenerateEncryptedPassword()
	if err := u.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var serializedUser models.UserEntity
	copier.Copy(&serializedUser, &user)
	ctx.JSON(http.StatusCreated, serializedUser)
}

// Login godoc
// @summary Login
// @tags auth
// @id Login
// @accept json
// @produce json
// @param Login body SignInRequest true "''"
// @router /auth/login [post]
func (u *User) SignIn(ctx *gin.Context) {
	var form SignInRequest
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var user models.UserEntity
	if check := u.DB.First(&user,"username = ?", form.Username); check.RowsAffected == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Not found username password"})
		return
	}

	if checkPassword := user.ValidateEncryptedPassword(form.Password); !checkPassword {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Not found username password"})
		return
	}

	accessToken , refreshToken := user.GenerateJwtToken()


	ctx.JSON(http.StatusOK, gin.H{"user":user,"accessToken" : accessToken,"refreshToken" :refreshToken})
}

// Refresh token godoc
// @summary Refresh token
// @tags auth
// @id Refresh token
// @produce plain
// @security BearerAuth
// @router /auth/refresh [get]
func (u *User) Refresh(ctx *gin.Context) {
	user := middleware.ExtractUserToken(ctx)
	accessToken , refreshToken := user.GenerateJwtToken()
	ctx.JSON(http.StatusOK, gin.H{"user":user,"accessToken" : accessToken,"refreshToken" :refreshToken})
}


