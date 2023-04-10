package handler

import (
	"nkrumahthis/assorted-jollof/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {

	users, err := repository.FindAllUsers()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("invalid user ID")
	}

	user, err := repository.FindUser(id)
	if err != nil {
		return c.Status(500).SendString("Failed to get user: " + err.Error())
	}

	return c.JSON(user)

}

func GetUsersByFilter(c *fiber.Ctx) error {
	// Parse the query parameters from the URL

	email := c.Query("email")
    name := c.Query("name")
    password := c.Query("password")

	users, err := repository.FindUsersWithFilter(name, email, password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	
	return c.JSON(users)
}
