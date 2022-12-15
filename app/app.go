package app

import (
	"log"

	sf_models "github.com/Subha-Research/koham/app/models"
	routes "github.com/Subha-Research/koham/app/routes/v1"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type KohamApp struct {
	role_model   sf_models.RoleModel
	access_model sf_models.AccessModel
}

func (k_app *KohamApp) SetupApp() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	routes.SetupPingRoute(app)

	database := sf_models.Database{}
	role_coll, _, err := database.GetCollectionAndSession("sf_roles")
	if err != nil {
		log.Fatal("Errro in  getting collection and session. Stopping server", err)
	}
	// Dependency injection pattern
	k_app.role_model.Collection = role_coll
	k_app.role_model.InsertAllRoles()

	access_coll, _, err := database.GetCollectionAndSession("sf_accesses")
	if err != nil {
		log.Fatal("Error in  getting collection and session. Stopping server", err)
	}
	// Dependency injection pattern
	k_app.access_model.Collection = access_coll
	k_app.access_model.InsertAllAccesses()

	routes.SetupRoutes(app)

	// Return configured app
	return app
}
