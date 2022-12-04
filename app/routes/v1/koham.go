package routes

import (
	base_controllers "github.com/Subha-Research/koham/app/base-controllers/v1"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	bc := base_controllers.BaseController{}
	v1 := api.Group("/:stakeholder_type/:user_type/:user_id/:resource_type")
	v1.Get("/", bc.GetHandler)
}
