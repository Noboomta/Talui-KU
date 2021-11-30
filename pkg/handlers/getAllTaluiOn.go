package handlers

import (
	_ "database/sql"
	"github.com/gofiber/fiber/v2"
	"talui/pkg/database/dbmodels"
)

func GetAllTaluiOnHandler(c *fiber.Ctx) error {
	taluiOn, err := dbmodels.GetAllTaluiOn()
	if err != nil {
		return c.JSON(fiber.Map{
			"msg": "error",
		})
	}
	
	return c.JSON(taluiOn)
}