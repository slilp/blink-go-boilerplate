package main

import (
	"blink-go-gin-boilerplate/config"
	"blink-go-gin-boilerplate/migration"
	"blink-go-gin-boilerplate/routes"
	"log"
	"os"

	_ "blink-go-gin-boilerplate/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Blink GO(Gin) Boilerplate
// @version         1.0.0

// @contact.name   Contact
// @contact.url    https://www.blink-me-code.dev/portfolio

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main(){
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error environment config")
	}

	db := config.InitDB()
	migration.Migrate(db)
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.Serve(r,db)

	r.Run(":"+os.Getenv("PORT"))

}