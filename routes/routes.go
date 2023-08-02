package routes

import (
	controllers "blink-go-gin-boilerplate/controllers"
	"blink-go-gin-boilerplate/middleware"
	"blink-go-gin-boilerplate/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Serve(r *gin.Engine , db *gorm.DB) {
	r.GET("/health-check",controllers.HealthCheck)

	authGroup := r.Group("/auth")
	userController := controllers.User{DB: db}
	{
		authGroup.POST("/login"  , userController.SignIn)
		authGroup.GET("/refresh" , middleware.RefreshUser() , userController.Refresh)
	}

	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", userController.Register)
	}

	productGroup := r.Group("/product")
	productController := controllers.Product{DB: db}
	{
		productGroup.POST("",middleware.AuthorizedUser([]models.RoleType{"ADMIN"}), productController.Create)
		productGroup.PUT("/:id", productController.Update)
		productGroup.DELETE("/:id",middleware.AuthorizedUser([]models.RoleType{"ADMIN"}), productController.Delete)
		productGroup.GET("/:id" , productController.FindOne)
	}

	orderGroup := r.Group("/order")
	orderController := controllers.Order{DB: db}
	{
		orderGroup.POST("",middleware.AuthorizedUser([]models.RoleType{"USER"}), orderController.Create)
		orderGroup.PATCH("/order-status/:id",middleware.AuthorizedUser([]models.RoleType{"ADMIN"}), orderController.UpdateStatus)
		orderGroup.PATCH("/:id",middleware.AuthorizedUser([]models.RoleType{"USER"}), orderController.Update)
		orderGroup.DELETE("/:id",middleware.AuthorizedUser([]models.RoleType{"USER"}), orderController.Delete)
		orderGroup.GET("/:id",middleware.AuthorizedUser([]models.RoleType{"USER"}), orderController.FindOne)
		
	}

}
