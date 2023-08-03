package main

import (
	"blink-go-gin-boilerplate/config"
	"blink-go-gin-boilerplate/migration"
	"blink-go-gin-boilerplate/routes"
	"log"
	"net/http"
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

var (
	app *gin.Engine
)

func main(){
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error environment config")
	}

	db := config.InitDB()
	migration.Migrate(db)
	app := gin.Default()
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.Serve(app,db)

	app.Run(":"+os.Getenv("PORT"))

}

func Handler(w http.ResponseWriter, r *http.Request){
	app.ServeHTTP(w,r)
}