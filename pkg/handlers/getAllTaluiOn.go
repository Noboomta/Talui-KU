package handlers

import (
	_ "database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"talui/pkg/database/dbmodels"
)

func GetAllTaluiOnHandler(c *fiber.Ctx) error {
	fmt.Println("/on/getAllTalui")
	
	taluiOn, err := dbmodels.GetAllTaluiOn()
	if err != nil {
		fmt.Println("[/on/GetAllTalui] fail")
		return c.JSON(fiber.Map{
			"msg": "error",
		})
	}
	
	return c.JSON(taluiOn)
}