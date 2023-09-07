package product

import (
	product "github.com/slilp/blink-go-boilerplate/app/product/api"
	"github.com/slilp/blink-go-boilerplate/app/product/internal"
	"github.com/slilp/blink-go-boilerplate/app/product/server"

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