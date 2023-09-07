package order

import (
	order "github.com/slilp/blink-go-boilerplate/app/order/api"
	"github.com/slilp/blink-go-boilerplate/app/order/internal"
	"github.com/slilp/blink-go-boilerplate/app/order/server"

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