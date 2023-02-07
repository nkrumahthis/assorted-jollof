package controller

import (
	"database/sql"
	"nkrumahthis/assorted-jollof/model"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	// Connect to the database
	db, err := sql.Open("sqlite3", "../main.db")
	if err != nil {
		return c.Status(500).SendString("Could not connect to the database")
	}
	defer db.Close()

	users, err := model.GetUsersFromDB(db)
	if err != nil {
		return c.Status(500).SendString("Could not retrieve users")
	}

	return c.JSON(users)
}