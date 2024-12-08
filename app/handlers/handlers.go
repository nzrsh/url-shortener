package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/nzrsh/url-shortener/app/services"
)

func SetupRouter(app *fiber.App) {

	app.Get("/home", HomePage)

	app.Post("/create", PostUrl)

	app.Get("/short/:short", func(c fiber.Ctx) error {
		shortUrl := c.Params("short")
		log.Info(shortUrl)
		source, err := services.GetLongUrl(shortUrl)
		if source == "no" && source == "nil" {
			c.SendString(err.Error())
		}

		return c.Redirect().Status(fiber.StatusSeeOther).To(source)
	})
}

func HomePage(c fiber.Ctx) error {
	return c.SendFile("./app/static/html/index.html")
}

func PostUrl(c fiber.Ctx) error {
	url := c.Query("url")
	newLink := services.CreateShortURL(url)
	return c.SendString(newLink.ShortLink)
}
