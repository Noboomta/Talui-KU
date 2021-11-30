package main

import (
	_ "log"
	"os"
	"talui/pkg/app"
	"talui/pkg/database"
	"talui/pkg/route"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv"
)

func main() {
	fiberApp := app.CreateFiberApp()
	
	// err := godotenv.Load()
    // if err != nil {
    //     log.Fatal("Error loading .env file")
    // }
    
    mysqlUri := os.Getenv("MYSQL_URI")
	database.ConnectMySQL(mysqlUri)
	// middleware.RegisDb(fiberApp, db)
	route.SetupRouter(fiberApp)
	
	port := os.Getenv("PORT")
	fiberApp.Listen(":" + port)
}