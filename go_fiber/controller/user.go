package controller

import (
	"nkrumahthis/assorted-jollof/database"
	"nkrumahthis/assorted-jollof/model"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	// Connect to the database
	db, err := database.GetDB()

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer db.Close()

	users, err := model.GetUsersFromDB(db)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(users)
}