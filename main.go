package main

import (
	"database/sql"
	"fmt"
	"os"

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

func ConnectMySQL() *sql.DB {
	db, err := sql.Open("mysql", "b6210545734:puvana.s@ku.th@tcp(iot.cpe.ku.ac.th:3306)/b6210545734")
	
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	
	return db
}

func main() {
	fiberApp := CreateFiberApp()
	db := ConnectMySQL()
	
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