package app

import (
	"blink-go-gin-boilerplate/app/order"
	"blink-go-gin-boilerplate/app/product"
	user "blink-go-gin-boilerplate/app/user"
	"blink-go-gin-boilerplate/database"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "blink-go-gin-boilerplate/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var server *http.Server
var router *gin.Engine

func Initialize() {
	config, err := loadConfig()

	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

	router := InitGin(*config)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		  "message": "health check",
		})
	  })

	db, err := database.Initialize(config.Database)
	if err != nil {
		panic(fmt.Errorf("failed to create database connection: %w", err))
	}

	user.Initialize(db, &router.RouterGroup) 
	product.Initialize(db, &router.RouterGroup)
	order.Initialize(db, &router.RouterGroup)


	Run(*config)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	Shutdown(ctx)

}


func InitGin(config Config) *gin.Engine {

	if config.Mode == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router = gin.New()
	router.Use(
		gin.Recovery(),
	)

	return router
}

func Run(config Config) *http.Server{
	
	server = &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: router,
	}
	fmt.Printf("Server is running on port: %s\n", server.Addr)

	go func() {
		if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Fatal("ListenAndServe")
		}
	}()

	return server
}

func Shutdown(ctx context.Context) {
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown")
	}
}