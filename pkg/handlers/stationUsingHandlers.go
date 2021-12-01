package handlers

import (
	_ "database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"talui/pkg/database/dbmodels"
)

func StationUsingEntryHandler(c *fiber.Ctx) error {
	fmt.Println("/station/using/entry")
	
	stationUsingData, err := dbmodels.GetStationUsingEntry()
	if err != nil {
		fmt.Println("[/station/using/entry] fail")
		return c.JSON(fiber.Map{
			"msg": "error",
		})
	}
	
	return c.JSON(stationUsingData)
}

func StationUsingDestHandler(c *fiber.Ctx) error {
	fmt.Println("/station/using/dest")
	
	stationUsingData, err := dbmodels.GetStationUsingDest()
	if err != nil {
		fmt.Println("[/station/using/dest] fail")
		return c.JSON(fiber.Map{
			"msg": "error",
		})
	}
	
	return c.JSON(stationUsingData)
}

func FindStationEntryHandler(c *fiber.Ctx) error {
	fmt.Println("/station/find/entry");
	
	station := c.Query("station");
	time := c.Query("time");
	
	stationUsingData, err := dbmodels.FindStationUsingEntryByStationAndTime(station, time);
	if err != nil {
		fmt.Println("[/station/find/entry] fail")
		return c.JSON(fiber.Map{
			"msg": "error",
		})
	}
	
	return c.JSON(stationUsingData)
}

func FindStationDestHandler(c *fiber.Ctx) error {
	fmt.Println("/station/find/dest");
	
	station := c.Query("station");
	time := c.Query("time");
	
	stationUsingData, err := dbmodels.FindStationUsingDestByStationAndTime(station, time);
	if err != nil {
		fmt.Println("[/station/find/dest] fail")
		return c.JSON(fiber.Map{
			"msg": "error",
		})
	}
	
	return c.JSON(stationUsingData)
}
