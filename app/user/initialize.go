package product

import (
	user "github.com/slilp/blink-go-boilerplate/app/user/api"
	"github.com/slilp/blink-go-boilerplate/app/user/internal"
	"github.com/slilp/blink-go-boilerplate/app/user/server"

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