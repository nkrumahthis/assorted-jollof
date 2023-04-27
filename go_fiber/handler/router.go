package handler

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/users", GetUsersByFilter)
	app.Get("/users/:id", GetUser)

	app.Get("/orders", GetOrdersByFilter)
	app.Get("/orders/:id", GetOrder)

	app.Get("/customers", GetCustomers)
	app.Get("/customers/:id", GetCustomer)
}
