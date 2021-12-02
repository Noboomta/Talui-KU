package handlers

import (
	_ "database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"talui/pkg/database/dbmodels"
)

// ShowAllTaluiComplete godoc
// @Summary Show all TaluiComplete
// @Description get all Talui
// @Accept  json
// @Produce  json
// @Success 200 {object} json
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router //complete/GetAllTalui [get]
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