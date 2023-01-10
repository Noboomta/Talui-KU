package handlers

import (
	_ "database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"talui/pkg/database/dbmodels"
)

func InsertTaluiCompleteHandler(c *fiber.Ctx) error {
	fmt.Println("/complete/insertTalui")

	entry := c.Query("entry")
	entry_ts := c.Query("entry_ts")
	dest := c.Query("dest")
	line := c.Query("line")
	
	err := dbmodels.InsertTaluiComplete(entry, entry_ts, dest, line)
	if err != nil {
		fmt.Printf("[TaluiComplete] Error Insert %s->%s on(%s) line: %s", entry, entry_ts, dest, line)
		return c.SendStatus(400)
	}
	fmt.Printf("[TaluiComplete] Complete Insert %s->%s on(%s) line: %s", entry, entry_ts, dest, line)
	return c.SendStatus(200)
}
