package router

import (
	"nkrumahthis/assorted-jollof/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App){
	app.Get("/users", controller.GetUsers)
}