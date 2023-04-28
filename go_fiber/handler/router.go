package handler

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/users", GetUsersByFilter)
	app.Get("/users/:id", GetUser)

	app.Get("/orders", GetOrdersByFilter)
	app.Get("/orders/:id", GetOrder)

	app.Get("/customers", GetCustomersByFilter)
	app.Get("/customers/:id", GetCustomer)

	app.Post("/payments", CreatePayment)
	app.Get("/payments", GetPaymentsByFilter)

}
