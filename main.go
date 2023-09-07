package main

import "github.com/slilp/blink-go-boilerplate/app"

// @title           Blink GO(Gin) Boilerplate
// @version         2.0.0

// @contact.name   Contact
// @contact.url    https://www.blink-me-code.dev/portfolio

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	app.Initialize()
}