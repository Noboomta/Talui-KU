package handlers

import (
	_ "database/sql"
	"fmt"
	"talui/pkg/database/dbmodels"

	"github.com/gofiber/fiber/v2"
)

func GetAllTaluiTrackerHandler(c *fiber.Ctx) error {
	fmt.Println("/GetAllTaluiTrackerHandler")

	taluiOn, err := dbmodels.GetTaluiTracker()
	if err != nil {
		fmt.Println("[/GetAllTaluiTrackerHandler] fail")
		return c.JSON(fiber.Map{
			"msg": "error",
		})
	}

	return c.JSON(taluiOn)
}
