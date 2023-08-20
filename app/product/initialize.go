package product

import (
	product "blink-go-gin-boilerplate/app/product/api"
	"blink-go-gin-boilerplate/app/product/internal"
	"blink-go-gin-boilerplate/app/product/server"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Initialize(
	db *gorm.DB,
	group *gin.RouterGroup,
) product.Service {
	repo := internal.NewRepository(db)
	service := internal.NewService(
		repo,
	)
	server.HttpMount(group, service)
	return service
}