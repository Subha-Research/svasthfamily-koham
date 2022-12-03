package app

import (
	"log"

	routes "github.com/Subha-Research/koham/app/routes/v1"
	sf_models "github.com/Subha-Research/koham/app/svasthfamily/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupApp() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	routes.SetupPingRoute(app)

	database := sf_models.Database{}
	collection, _, err := database.GetCollectionAndSession("sf_roles")
	log.Println(collection, err)

	rm := sf_models.RoleModel{}
	rm.InsertAllRoles(collection)

	routes.SetupRoutes(app)

	// Return configured app
	return app
}
