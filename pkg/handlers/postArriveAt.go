package handlers

import (
	_ "database/sql"
	"talui/pkg/database/dbmodels"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func ArriveAtHandler(c *fiber.Ctx) error {
	fmt.Println("/ArriveAt")
	
	dest := c.Query("dest")
	line := c.Query("line")
	
	// find by line, dest in talui_on
	dataArray, err := dbmodels.FindManyTaluiOn(dest, line)
	if err != nil {
		fmt.Println("[ArriveAt] findMany fail")
		return c.SendStatus(400)
	}
	
	// InsertTaluiComplete for each talui to talui_complete
	for _, talui := range dataArray {
		err := dbmodels.InsertTaluiComplete(talui.Entry, talui.Entry_ts, talui.Dest, talui.Line)
		if err != nil {
			fmt.Println("[ArriveAt] InsertTaluiComplete fail")
			return c.SendStatus(400)
		}
	}
	
	// remove data in talui_on
	_, err2 := dbmodels.RemoveManyTaluiOn(dest, line)
	if err2 != nil {
		fmt.Println("[ArriveAt] removeMany fail")
		return c.SendStatus(400)
	}
	
	return c.SendStatus(200)
}