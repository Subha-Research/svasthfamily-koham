package routes

import (
	sf_controllers "github.com/Subha-Research/koham/app/controllers/v1"
	"github.com/gofiber/fiber/v2"
)

func SetupPingRoute(app *fiber.App) {
	api := app.Group("/ping")
	api.Post("/", sf_controllers.PingHandler)
	api.Get("/", sf_controllers.PingGetHandler)
}
