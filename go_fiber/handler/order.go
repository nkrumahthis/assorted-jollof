package handler

import (
	"nkrumahthis/assorted-jollof/model"

	"github.com/gofiber/fiber/v2"
)

func GetOrders(c *fiber.Ctx) error {
	orders, err := model.FindAll()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}	
	return c.JSON(orders)
}