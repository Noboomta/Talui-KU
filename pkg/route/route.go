package route

import (
	"github.com/gofiber/fiber/v2"
	"talui/pkg/handlers"
)

func SetupRouter(app *fiber.App) {
	app.Get("/on/getAllTalui", handlers.GetAllTaluiOnHandler)
	app.Post("/on/insertTalui", handlers.InsertTaluiOnHandler)
	app.Get("/complete/GetAllTalui", handlers.GetAllTaluiCompleteHandler)
	app.Post("/complete/insertTalui", handlers.InsertTaluiCompleteHandler)
	app.Post("/arriveAt", handlers.ArriveAtHandler)
	app.Get("/", handlers.GetMainPage)
}
