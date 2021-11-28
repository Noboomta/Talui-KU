package main

import (
	"database/sql"
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

type Talui struct {
	id                                      int
	enter, down, enter_time, down_time     string
}

func CreateFiberApp() *fiber.App {
	app := fiber.New()
	return app
}

func ConnectMySQL(uri string) *sql.DB {
	fmt.Println(uri)
	db, err := sql.Open("mysql", uri)
	
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	
	return db
}

func main() {
	fiberApp := CreateFiberApp()
	
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    
    mysql_uri := os.Getenv("MYSQL_URI")
	db := ConnectMySQL(mysql_uri)
	
	
	fiberApp.Get("/test", func(c *fiber.Ctx) error {
		fmt.Println("someone using /test")
		var result = "hello"
		return c.JSON(result)
	})
	
	fiberApp.Post("/insert", func(c *fiber.Ctx) error {
		insert, err := db.Query("INSERT INTO talui_using (enter, down) VALUES ( 'staA', 'staB' )")
		
		if err != nil {
			fmt.Println(err)
		}
		defer insert.Close()
		
		return c.JSON("done")
	}) 
	
	fiberApp.Get("/get", func(c *fiber.Ctx) error {
		results, err := db.Query("SELECT * FROM talui_using")
		
		if err != nil {
			fmt.Println(err)
		}
		defer results.Close()
		
		for results.Next() {
			var talui Talui
			results.Scan(&talui.id, &talui.enter, &talui.enter_time, &talui.down, &talui.down_time)
			fmt.Println(talui.id, talui.enter, talui.enter_time, talui.down, talui.down_time)
		}
		
		return c.JSON("done")
	})
	
	fiberApp.Listen(":3000")
}