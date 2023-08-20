package server

import (
	user "blink-go-gin-boilerplate/app/user/api"
	"blink-go-gin-boilerplate/middleware"
	"blink-go-gin-boilerplate/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func HttpMount(router gin.IRouter, service user.Service) {
	handler := NewHttpHandler(service) 
	userGroup := router.Group("user")
	{
		userGroup.POST("/register", handler.register)
	}

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login"  , handler.signin)
		authGroup.GET("/refresh" , middleware.RefreshUser() , handler.refresh)
	}


}

func NewHttpHandler(service user.Service) *handler {
	return &handler{service: service}
}

type handler struct {
	service user.Service
}

// Register godoc
// @summary Register
// @tags user
// @id Register
// @accept json
// @produce json
// @param User body user.RegisterRequest true "''"
// @router /user/register [post]
func (h *handler) register(c *gin.Context) {
	ctx := utils.HttpStatusContext{Context: c}

	var form user.RegisterRequest
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.BadRequest(utils.ERROR_INVALID_REQUEST)
		return
	}

	if userInfo , _ := h.service.GetByUsername(form.Username); userInfo.Username != "" {
		ctx.BadRequest(utils.ERROR_ALREADY_USED_USER)
		return
	}


	var user user.UserEntity
	copier.Copy(&user, &form)

	user.Password = utils.GenerateEncryptedPassword(user.Password)
	if  _ ,err := h.service.Create(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": utils.ERROR_CREATE_USER})
		return
	}

	
	ctx.Created(nil)
}

// Login godoc
// @summary Login
// @tags auth
// @id Login
// @accept json
// @produce json
// @param Login body user.SignInRequest true "''"
// @router /auth/login [post]
func (h *handler) signin(c *gin.Context) {
	ctx := utils.HttpStatusContext{Context: c}

	var form user.SignInRequest
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.BadRequest(utils.ERROR_INVALID_REQUEST)
		return
	}
	userInfo , _ := h.service.GetByUsername(form.Username)
	if   userInfo.Username == "" {
		ctx.BadRequest(utils.ERROR_USERNAME_PASSWORD)
		return
	}

	if checkPassword := utils.ValidateEncryptedPassword(form.Password,userInfo.Password); !checkPassword {
		ctx.BadRequest(utils.ERROR_USERNAME_PASSWORD)
		return
	}

	accessToken , refreshToken := utils.GenerateJwtToken(userInfo)


	ctx.OK(gin.H{"user":userInfo,"accessToken" : accessToken,"refreshToken" :refreshToken})
}

// Refresh token godoc
// @summary Refresh token
// @tags auth
// @id Refresh token
// @produce plain
// @security BearerAuth
// @router /auth/refresh [get]
func (h *handler) refresh(c *gin.Context) {
	ctx := utils.HttpStatusContext{Context: c}

	user := middleware.ExtractUserToken(c)
	accessToken , refreshToken := utils.GenerateJwtToken(user)
	ctx.OK(gin.H{"user":user,"accessToken" : accessToken,"refreshToken" :refreshToken})
}

