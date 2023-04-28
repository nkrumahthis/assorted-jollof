package handler

import (
	"nkrumahthis/assorted-jollof/repository"

	"github.com/gofiber/fiber/v2"
)

func CreatePayment(c *fiber.Ctx) error {
	payload := repository.Payment{}

	err := c.BodyParser(&payload)

	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON(payload)
}

func GetPaymentsByFilter(c *fiber.Ctx) error {
	// collect query params from request
	id := c.Query("id")
	amount := c.Query("amount")
	order_id := c.Query("order_query")
	timestamp := c.Query("timestamp")
	user_id := c.Query("user_id")

	payments, err := repository.FindPaymentsByParams(id, amount, order_id, timestamp, user_id)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}
	return c.JSON(payments)
}
