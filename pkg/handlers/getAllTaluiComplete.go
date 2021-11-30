package handlers

import (
	_ "database/sql"
	"github.com/gofiber/fiber/v2"
	"talui/pkg/database/dbmodels"
)

func GetAllTaluiCompleteHandler(c *fiber.Ctx) error {
	taluiOn, err := dbmodels.GetAllTaluiComplete()
	if err != nil {
		return c.JSON(fiber.Map{
			"msg": "error",
		})
	}
	
	return c.JSON(taluiOn)
}