package routes

import (
	controllers "github.com/Subha-Research/koham/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupPingRoute(app *fiber.App) {
	api := app.Group("/ping")
	api.Post("/", controllers.PingHandler)
	api.Get("/", controllers.PingGetHandler)
}
