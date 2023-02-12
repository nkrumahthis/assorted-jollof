package handler

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/users", GetUsers)
	app.Get("/users/:id", GetUser)

	app.Get("/orders", GetOrders)
	app.Get("/orders/:id", GetOrder)

	app.Get("/customers", GetCustomers)
	app.Get("/customers/:id", GetCustomer)
}
