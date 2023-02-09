package router

import (
	"nkrumahthis/assorted-jollof/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App){
	app.Get("/users", handler.GetUsers)
}