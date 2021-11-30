package handlers

import (
	_ "database/sql"
	"talui/pkg/database/dbmodels"

	"github.com/gofiber/fiber/v2"
)

func ArriveAtHandler(c *fiber.Ctx) error {
	dest := c.Query("dest")
	line := c.Query("line")
	
	// find by line, dest in talui_on
	dataArray, err := dbmodels.FindManyTaluiOn(dest, line)
	if err != nil {
		return c.SendStatus(400)
	}
	
	// InsertTaluiComplete for each talui to talui_complete
	for _, talui := range dataArray {
		dbmodels.InsertTaluiComplete(talui.Entry, talui.Entry_ts, talui.Dest, talui.Line)
	}
	
	// remove data in talui_on
	_, err2 := dbmodels.RemoveManyTaluiOn(dest, line)
	if err2 != nil {
		return c.SendStatus(400)
	}
	
	return c.SendStatus(200)
}