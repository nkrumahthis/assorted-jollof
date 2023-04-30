package handler

import (
	"nkrumahthis/assorted-jollof/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateCustomer(c *fiber.Ctx) error {
	// expected payload
	payload := struct {
		Name  string
		Phone string
		Token string
	}{}

	// parsing payload
	err := c.BodyParser(&payload)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// create new customer in repository
	customer, err := repository.CreateCustomer(payload.Name, payload.Phone, payload.Token)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}
	return c.JSON(customer)
}

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
