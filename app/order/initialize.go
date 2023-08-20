package order

import (
	order "blink-go-gin-boilerplate/app/order/api"
	"blink-go-gin-boilerplate/app/order/internal"
	"blink-go-gin-boilerplate/app/order/server"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Initialize(
	db *gorm.DB,
	group *gin.RouterGroup,
) order.Service {
	repo := internal.NewRepository(db)
	service := internal.NewService(
		repo,
	)
	server.HttpMount(group, service)
	return service
}