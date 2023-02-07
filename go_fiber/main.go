package main

import (
	"nkrumahthis/assorted-jollof/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Assorted Jollof!")
	})

	router.SetupRoutes(app)

	app.Listen(":3000")
}