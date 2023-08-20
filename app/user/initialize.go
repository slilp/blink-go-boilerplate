package product

import (
	user "blink-go-gin-boilerplate/app/user/api"
	"blink-go-gin-boilerplate/app/user/internal"
	"blink-go-gin-boilerplate/app/user/server"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Initialize(
	db *gorm.DB,
	group *gin.RouterGroup,
) user.Service {
	repo := internal.NewRepository(db)
	service := internal.NewService(
		repo,
	)
	server.HttpMount(group, service)
	return service
}