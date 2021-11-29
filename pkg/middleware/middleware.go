package middleware

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

func RegisDb(app *fiber.App, db *sql.DB){
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("user", "admin")
		c.Locals("db", db)
		return c.Next()
	})
}