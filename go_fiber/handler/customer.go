package handler

import (
	"nkrumahthis/assorted-jollof/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetCustomers(c *fiber.Ctx) error {
	customers, err := repository.FindAllCustomers()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(customers)

}

func GetCustomer(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(500).SendString("invalid customer id")
	}
	customers, err := repository.FindCustomer(id)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(customers)

}

func GetCustomersByFilter(c *fiber.Ctx) error {
	// collect query params from request
	id := c.Query("id")
	name := c.Query("name")
	phone := c.Query("phone")
	token := c.Query("token")

	// find customers by those params
	customers, err := repository.FindCustomersByParams(id, name, phone, token)

	// return response from query
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(customers)
}