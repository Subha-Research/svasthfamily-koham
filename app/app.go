package app

import (
	"log"

	routes "github.com/Subha-Research/koham/app/routes/v1"
	sf_models "github.com/Subha-Research/koham/app/svasthfamily/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type KohamApp struct {
	role_model sf_models.RoleModel
	// access_model sf_models.AccessModel
}

func (k_app *KohamApp) SetupApp() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	routes.SetupPingRoute(app)

	database := sf_models.Database{}
	collection, _, err := database.GetCollectionAndSession("sf_roles")
	log.Println(collection, err)
	// Dependency injection pattern
	k_app.role_model.Collection = collection
	k_app.role_model.InsertAllRoles()
	routes.SetupRoutes(app)

	// Return configured app
	return app
}
