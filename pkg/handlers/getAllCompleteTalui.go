package handlers

import (
	_ "database/sql"
	"github.com/gofiber/fiber/v2"
	"talui/pkg/database/dbmodels"
)

func GetAllCompleteTalui(c *fiber.Ctx) error {
	taluis, err := dbmodels.GetAllTalui()
	if err != nil {
		return c.JSON(fiber.Map{
			"msg": "error",
		})
	}
	
	return c.JSON(taluis)
}