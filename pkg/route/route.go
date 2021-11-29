package route

import (
	"github.com/gofiber/fiber/v2"
	"talui/pkg/handlers"
)

func SetupRouter(app *fiber.App) {
	app.Get("/all_taluis", handlers.GetAllCompleteTalui)
}
