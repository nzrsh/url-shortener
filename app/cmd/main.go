package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/nzrsh/url-shortener/app/handlers"
	"github.com/nzrsh/url-shortener/app/middlewares"
	"github.com/nzrsh/url-shortener/app/services"
)

func main() {
	app := fiber.New()
	services.CreateOrConnectDB("./repos/db.db")
	middlewares.SetupMiddlewares(app)
	handlers.SetupRouter(app)
	err := app.Listen("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
