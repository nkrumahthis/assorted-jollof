package handler

import (
	"nkrumahthis/assorted-jollof/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetOrders(c *fiber.Ctx) error {
	orders, err := repository.FindAllOrders()
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