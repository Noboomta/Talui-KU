package handlers

import (
	_ "database/sql"
	"github.com/gofiber/fiber/v2"
	"talui/pkg/database/dbmodels"
)

func InsertTaluiCompleteHandler(c *fiber.Ctx) error {
	entry := c.Query("entry")
	entry_ts := c.Query("entry_ts")
	dest := c.Query("dest")
	line := c.Query("line")
	
	err := dbmodels.InsertTaluiComplete(entry, entry_ts, dest, line)
	if err != nil {
		return c.SendStatus(400)
	}
	
	return c.SendStatus(200)
}