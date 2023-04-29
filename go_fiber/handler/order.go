package handler

import (
	"nkrumahthis/assorted-jollof/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateOrder(c *fiber.Ctx) error {
	payload := struct {
		Packs      int
		CustomerId int
		Location   string
		Status     string
	}{}

	// load data from body of request
	err := c.BodyParser(&payload)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	if payload.Status == "" {
		payload.Status = "PENDING"
	}

	// create a new order
	order, err := repository.CreateOrder(payload.Packs, payload.CustomerId, payload.Location, payload.Status)

	// if error, send message
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	// respond with order
	return c.JSON(order)
}

func GetOrders(c *fiber.Ctx) error {
	orders, err := repository.FindAllOrders()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(orders)
}

func GetOrdersByFilter(c *fiber.Ctx) error {
	id := c.Query("id")
	packs := c.Query("packs")
	customerId := c.Query("customer_id")
	location := c.Query("location")
	status := c.Query("status")
	createdAt := c.Query("created_at")

	orders, err := repository.FindOrdersByParams(id, packs, customerId, location, status, createdAt)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(orders)
}

func GetOrder(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(500).SendString("invalid order id")
	}
	order, err := repository.FindOrder(id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(order)
}
