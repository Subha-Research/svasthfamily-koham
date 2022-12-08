package app

import (
	"log"

	routes "github.com/Subha-Research/koham/app/routes/v1"
	sf_models "github.com/Subha-Research/koham/app/svasthfamily/models"
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
	log.Println(role_coll, err)
	// Dependency injection pattern
	k_app.role_model.Collection = role_coll
	k_app.role_model.InsertAllRoles()

	access_coll, _, err := database.GetCollectionAndSession("sf_accesses")
	log.Println(access_coll, err)
	// Dependency injection pattern
	k_app.access_model.Collection = access_coll
	k_app.access_model.InsertAllAccesses()

	routes.SetupRoutes(app)

	// Return configured app
	return app
}
