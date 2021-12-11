package route

import (
	"talui/pkg/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	app.Get("/getTracker", handlers.GetAllTaluiTrackerHandler)
	app.Get("/on/getAllTalui", handlers.GetAllTaluiOnHandler)
	app.Post("/on/insertTalui", handlers.InsertTaluiOnHandler)
	app.Get("/complete/GetAllTalui", handlers.GetAllTaluiCompleteHandler)
	app.Post("/complete/insertTalui", handlers.InsertTaluiCompleteHandler)
	app.Post("/arriveAt", handlers.ArriveAtHandler)
	app.Get("/station/using/entry", handlers.StationUsingEntryHandler)
	app.Get("/station/using/dest", handlers.StationUsingDestHandler)
	app.Get("/station/find/entry", handlers.FindStationEntryHandler)
	app.Get("/station/find/dest", handlers.FindStationDestHandler)
	app.Get("/", handlers.GetMainPage)
}
