package tests

import (
	"io"
	"net/http"
	"testing"

	app "github.com/Subha-Research/svasthfamily-koham/app"
	base_validators "github.com/Subha-Research/svasthfamily-koham/app/base-validators"
	"github.com/Subha-Research/svasthfamily-koham/app/controllers/v1"
	"github.com/Subha-Research/svasthfamily-koham/app/models"
	"github.com/Subha-Research/svasthfamily-koham/app/routes/v1"
	"github.com/Subha-Research/svasthfamily-koham/app/services/v1"
	"github.com/Subha-Research/svasthfamily-koham/app/validators"
	models_mock "github.com/Subha-Research/svasthfamily-koham/tests/mocks/models"
	services_mock "github.com/Subha-Research/svasthfamily-koham/tests/mocks/services"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	tests := []struct {
		description string
		route       string
		// Expected output
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "ping route testcase",
			route:         "/ping",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "Ping is working.",
		},
	}

	f_app := app.InitFiberApplication()
	app := &app.KohamApp{
		App: f_app,
		Routes: &routes.Routes{
			BaseValidator: &base_validators.BaseValidator{
				ITokenService: &services.TokenService{},
			},
			BaseController: &controllers.BaseController{
				ACLController: &controllers.ACLController{
					Validator: &validators.ACLValidator{},
					Service: &services.ACLService{
						Model: &models.AccessRelationshipModel{},
					},
				},
				TokenController: &controllers.TokenController{
					Validator: &validators.TokenValidator{},
					IService: &services_mock.TokenServiceTest{
						Model:   &models_mock.TokenModelMock{},
						ARModel: &models_mock.AccessRelationshipModelMock{},
					},
				},
			},
		},

		DB:          &models.Database{},
		RoleModel:   &models.RoleModel{},
		AccessModel: &models.AccessModel{},
	}
	koham_app := app.SetupApp()
	// Iterate through testcases
	for _, test := range tests {
		req, _ := http.NewRequest(
			"GET",
			test.route,
			nil,
		)

		// The -1 disables request latency.
		res, _ := koham_app.Test(req, -1)

		// Verify if the status code is as expected
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)
		// Read the response body
		body, _ := io.ReadAll(res.Body)

		// Verify, that the reponse body equals the expected body
		assert.Equalf(t, test.expectedBody, string(body), test.description)
	}
}
