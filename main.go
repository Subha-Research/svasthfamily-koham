package main

import (
	"log"

	app "github.com/Subha-Research/svasthfamily-koham/app"
	base_validators "github.com/Subha-Research/svasthfamily-koham/app/base-validators"
	"github.com/Subha-Research/svasthfamily-koham/app/cache"
	"github.com/Subha-Research/svasthfamily-koham/app/constants"
	"github.com/Subha-Research/svasthfamily-koham/app/controllers/v1"
	"github.com/Subha-Research/svasthfamily-koham/app/models"
	"github.com/Subha-Research/svasthfamily-koham/app/routes/v1"
	"github.com/Subha-Research/svasthfamily-koham/app/services/v1"
	"github.com/Subha-Research/svasthfamily-koham/app/validators"
)

func main() {
	db := &models.Database{}
	token_coll, _, err := db.GetCollectionAndSession(constants.TOKEN_COLLECTION)
	if err != nil {
		log.Fatal("Error in getting collection and session. Stopping server", err)
	}
	acl_coll, _, err := db.GetCollectionAndSession(constants.ACL_COLLECTION)
	if err != nil {
		log.Fatal("Error in getting collection and session. Stopping server", err)
	}
	role_coll, _, err := db.GetCollectionAndSession(constants.ROLE_COLLECTION)
	if err != nil {
		log.Fatal("Error in getting collection and session. Stopping server", err)
	}
	access_coll, _, err := db.GetCollectionAndSession(constants.ACCESS_COLLECTION)
	if err != nil {
		log.Fatal("Error in getting collection and session. Stopping server", err)
	}

	f_app := app.InitFiberApplication()
	rc := cache.Redis{}
	token_rc := rc.SetupTokenRedisDB()
	app := app.KohamApp{}

	// Initialize all the dependencies
	app.Routes = &routes.Routes{
		BaseValidator: &base_validators.BaseValidator{
			ITokenService: &services.TokenService{
				Model: &models.TokenModel{
					Collection: token_coll,
				},
			},
		},
		BaseController: &controllers.BaseController{
			ACLController: &controllers.ACLController{
				Validator: &validators.ACLValidator{},
				Service: &services.ACLService{
					Model: &models.AccessRelationshipModel{
						Collection: acl_coll,
					},
				},
				ITokenService: &services.TokenService{
					Model: &models.TokenModel{
						Collection: token_coll,
					},
					ARModel: &models.AccessRelationshipModel{
						Collection: acl_coll,
					},
					Cache: &cache.TokenCache{
						RedisClient: &token_rc,
					},
				},
			},
			TokenController: &controllers.TokenController{
				Validator: &validators.TokenValidator{},
				IService: &services.TokenService{
					Model: &models.TokenModel{
						Collection: token_coll,
					},
					ARModel: &models.AccessRelationshipModel{
						Collection: acl_coll,
					},
					Cache: &cache.TokenCache{
						RedisClient: &token_rc,
					},
				},
			},
		},
	}
	app.RoleModel = &models.RoleModel{
		Collection: role_coll,
	}
	app.AccessModel = &models.AccessModel{
		Collection: access_coll,
	}
	app.App = f_app
	fiber_app := app.SetupApp()
	fiber_app.Listen(":8080")
}
