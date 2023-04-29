package handler

import (
	"nkrumahthis/assorted-jollof/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreatePayment(c *fiber.Ctx) error {
	payload := repository.Payment{}

	err := c.BodyParser(&payload)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	payment, err := repository.CreatePayment(payload)

	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON(payment)
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

func GetPayment(c *fiber.Ctx) error {
	// get id from query parameter
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON("invalid payment id")
	}

	// Find payment with that id
	payment, err := repository.FindPayment(id)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON(payment)
}
