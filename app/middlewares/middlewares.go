package middlewares

import (
	"github.com/gofiber/fiber/v3"
	_ "github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func SetupMiddlewares(app *fiber.App) error {
	app.Use("/static", static.New("./app/static"))
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	return nil
}
