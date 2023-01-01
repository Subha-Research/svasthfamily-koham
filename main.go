package main

import (
	app "github.com/Subha-Research/svasthfamily-koham/app"
	base_validators "github.com/Subha-Research/svasthfamily-koham/app/base-validators"
	"github.com/Subha-Research/svasthfamily-koham/app/cache"
	"github.com/Subha-Research/svasthfamily-koham/app/controllers/v1"
	"github.com/Subha-Research/svasthfamily-koham/app/models"
	"github.com/Subha-Research/svasthfamily-koham/app/routes/v1"
	"github.com/Subha-Research/svasthfamily-koham/app/services/v1"
	"github.com/Subha-Research/svasthfamily-koham/app/validators"
)

func main() {
	f_app := app.InitFiberApplication()
	rc := cache.Redis{}
	token_rc := rc.SetupTokenRedisDB()
	app := app.KohamApp{}

	// Initialize all the dependencies
	app.Routes = &routes.Routes{
		BaseValidator: &base_validators.BaseValidator{
			TokenService: &services.TokenService{},
		},
		BaseController: &controllers.BaseController{
			ACLController: &controllers.ACLController{
				Validator: validators.ACLValidator{},
				Service: services.ACLService{
					Model: &models.AccessRelationshipModel{},
				},
			},
			TokenController: &controllers.TokenController{
				Validator: &validators.TokenValidator{},
				Service: &services.TokenService{
					Model:   &models.TokenModel{},
					ARModel: &models.AccessRelationshipModel{},
					Cache: &cache.TokenCache{
						RedisClient: &token_rc,
					},
				},
			},
		},
	}
	app.DB = &models.Database{}
	app.RoleModel = &models.RoleModel{}
	app.AccessModel = &models.AccessModel{}
	app.App = f_app

	fiber_app := app.SetupApp()
	fiber_app.Listen(":8080")
}
