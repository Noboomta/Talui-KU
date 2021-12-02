package main

import (
	_ "log"
	"os"
	"talui/pkg/app"
	"talui/pkg/database"
	"talui/pkg/route"
	
	"github.com/gofiber/fiber/v2/middleware/cors"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv"
)

func main() {
	fiberApp := app.CreateFiberApp()
	
	// err := godotenv.Load()
    // if err != nil {
    //     log.Fatal("Error loading .env file")
    // }
    // Default config
	fiberApp.Use(cors.New())
	// Or extend your config for customization
	fiberApp.Use(cors.New(cors.Config{
	    AllowOrigins:     "*",
	    AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
	    AllowHeaders:     "",
	}))
    
    mysqlUri := os.Getenv("MYSQL_URI")
	database.ConnectMySQL(mysqlUri)
	// middleware.RegisDb(fiberApp, db)
	route.SetupRouter(fiberApp)
	
	port := os.Getenv("PORT")
	fiberApp.Listen(":" + port)
}