package handlers

import (
	_ "database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"talui/pkg/database/dbmodels"
)

func InsertTaluiOnHandler(c *fiber.Ctx) error {
	fmt.Println("/on/insertTalui")
	
	entry := c.Query("entry")
	dest := c.Query("dest")
	line := c.Query("line")
	
	err := dbmodels.InsertTaluiOn(entry, dest, line)
	if err != nil {
		fmt.Printf("[TaluiOn] Error Insert %s->%s line: %s", entry, dest, line)
		return c.SendStatus(400)
	}
	
	fmt.Printf("[TaluiOn] Complete Insert %s->%s line: %s", entry, dest, line)
	return c.SendStatus(200)
}
