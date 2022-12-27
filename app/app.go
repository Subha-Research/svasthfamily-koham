package app

import (
	"log"

	base_validators "github.com/Subha-Research/svasthfamily-koham/app/base-validators"
	"github.com/Subha-Research/svasthfamily-koham/app/controllers/v1"
	models "github.com/Subha-Research/svasthfamily-koham/app/models"
	routes "github.com/Subha-Research/svasthfamily-koham/app/routes/v1"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Improvement, this can be part of main file
func InitFiberApplication() *fiber.App {
	fiber_app := fiber.New()
	return fiber_app
}

type KohamApp struct {
	RoleModel   models.RoleModel
	AccessModel models.AccessModel
	Routes      routes.Routes
	BC          controllers.BaseController
	BV          base_validators.BaseValidator
	App         *fiber.App
}

func (k_app *KohamApp) SetupApp() *fiber.App {
	k_app.App.Use(logger.New())
	routes.SetupPingRoute(k_app.App)

	database := models.Database{}
	role_coll, _, err := database.GetCollectionAndSession("roles")
	if err != nil {
		log.Fatal("Errro in  getting collection and session. Stopping server", err)
	}
	// Dependency injection pattern
	k_app.RoleModel.Collection = role_coll
	k_app.RoleModel.InsertAllRoles()

	access_coll, _, err := database.GetCollectionAndSession("accesses")
	if err != nil {
		log.Fatal("Error in  getting collection and session. Stopping server", err)
	}
	// Dependency injection pattern
	k_app.AccessModel.Collection = access_coll
	k_app.AccessModel.InsertAllAccesses()

	k_app.Routes.Controller = k_app.BC
	k_app.Routes.Validator = k_app.BV
	k_app.Routes.SetupRoutes(k_app.App)
	// Return configured app
	return k_app.App
}
