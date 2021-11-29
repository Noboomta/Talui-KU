package main

import (
	"log"
	"os"
	"talui/pkg/app"
	"talui/pkg/database"
	"talui/pkg/database/dbmodels"
	"talui/pkg/route"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	fiberApp := app.CreateFiberApp()
	
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    
    mysqlUri := os.Getenv("MYSQL_URI")
	database.ConnectMySQL(mysqlUri)
	// middleware.RegisDb(fiberApp, db)
	route.SetupRouter(fiberApp)
	
	fiberApp.Get("/allTaluis", func(c *fiber.Ctx) error {
		taluis, err := dbmodels.GetAllTalui()
		if err != nil {
			return c.JSON(fiber.Map{
				"msg": "error",
			})
		}
		
		return c.JSON(taluis)
	})
	
	fiberApp.Listen(":3000")
}