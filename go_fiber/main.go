package main

import (
	"nkrumahthis/assorted-jollof/database"
	"nkrumahthis/assorted-jollof/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.PrepDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Assorted Jollof!")
	})

	handler.SetupRoutes(app)

	app.Listen(":3000")
}