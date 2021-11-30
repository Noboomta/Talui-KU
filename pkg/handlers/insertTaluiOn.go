package handlers

import (
	_ "database/sql"
	"github.com/gofiber/fiber/v2"
	"talui/pkg/database/dbmodels"
)

func InsertTaluiOnHandler(c *fiber.Ctx) error {
	entry := c.Query("entry")
	dest := c.Query("dest")
	line := c.Query("line")
	
	err := dbmodels.InsertTaluiOn(entry, dest, line)
	if err != nil {
		return c.SendStatus(400)
	}
	
	return c.SendStatus(200)
}