package app

import (
	routes "github.com/Subha-Research/pariwar-koham/app/routes/v1"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupApp() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	routes.SetupPingRoute(app)
	// routes.SetupRoutes(app)

	// Return configured app
	return app
}
