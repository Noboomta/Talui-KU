package handlers

import (
	_ "database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"talui/pkg/database/dbmodels"
)

func GetAllTaluiCompleteHandler(c *fiber.Ctx) error {
	fmt.Println("/complete/GetAllTalui")
	
	taluiOn, err := dbmodels.GetAllTaluiComplete()
	if err != nil {
		fmt.Println("[/complete/GetAllTalui] fail")
		return c.JSON(fiber.Map{
			"msg": "error",
		})
	}
	
	return c.JSON(taluiOn)
}