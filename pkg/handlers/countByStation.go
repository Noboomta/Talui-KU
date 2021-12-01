package handlers

import (
	_ "database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"talui/pkg/database/dbmodels"
)

func StationHandler(c *fiber.Ctx) error {
	fmt.Println("/station/using")
	
	stationCount, err := dbmodels.GetStationUsing()
	if err != nil {
		fmt.Println("[/station/using] fail")
		return c.JSON(fiber.Map{
			"msg": "error",
		})
	}
	
	return c.JSON(stationCount)
}