package tests

import (
	"encoding/json"
	"log"
	"net/http"
	"testing"
	"time"

	app "github.com/Subha-Research/svasthfamily-koham/app"
	base_validators "github.com/Subha-Research/svasthfamily-koham/app/base-validators"
	"github.com/Subha-Research/svasthfamily-koham/app/constants"
	"github.com/Subha-Research/svasthfamily-koham/app/controllers/v1"
	"github.com/Subha-Research/svasthfamily-koham/app/dto"
	"github.com/Subha-Research/svasthfamily-koham/app/models"
	"github.com/Subha-Research/svasthfamily-koham/app/routes/v1"
	"github.com/Subha-Research/svasthfamily-koham/app/services/v1"
	"github.com/Subha-Research/svasthfamily-koham/app/validators"

	// services_mock "github.com/Subha-Research/svasthfamily-koham/tests/mocks/services"
	models_mock "github.com/Subha-Research/svasthfamily-koham/tests/mocks/models"
	services_mock "github.com/Subha-Research/svasthfamily-koham/tests/mocks/services"

	"github.com/stretchr/testify/assert"
)

func TestCreateToken(t *testing.T) {
	loc, _ := time.LoadLocation("Local")
	tests := []struct {
		testcode    string
		description string
		route       string
		// Expected output
		expectedError bool
		expectedCode  int
		expectedBody  interface{}
	}{
		{
			testcode:      "CREATE_TOKEN",
			description:   "Create token API testcase",
			route:         "/api/v1/family/users/8204a616-2131-4a64-97d0-ae3f2b9211be/tokens",
			expectedError: false,
			expectedCode:  201,
			expectedBody: &dto.CreateTokenResponse{
				TokenKey:     "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJGVXNlcklEIjoiODIwNGE2MTYtMjEzMS00YTY0LTk3ZDAtYWUzZjJiOTIxMWJlIiwiQWNjZXNzTGlzdCI6W3siY2hpbGRfbWVtYmVyX2lkIjoiODIwNGE2MTYtMjEzMS00YTY0LTk3ZDAtYWUzZjJiOTIxMWJlIiwiYWNjZXNzX2VudW1zIjpbMTAxLDEwMiwxMDMsMTA0LDEwNSwxMDYsMTA3LDEwOCwxMDldfV0sImlzcyI6InN2YXN0aGZhbWlseS1rb2hhbSIsImV4cCI6MTYxMDA0NDIwMCwiaWF0IjoxNjA5NDM5NDAwfQ.VJ2ln28qTYIQxFSUacByUFbVEDajg7v-inK-ySboQ78",
				TokenExpiry:  time.Date(2021, 1, 1, 0, 0, 0, 0, loc).Add(constants.TokenExpiryTTL * time.Hour),
				FamilyUserID: "8204a616-2131-4a64-97d0-ae3f2b9211be",
			},
		},
	}

	f_app := app.InitFiberApplication()
	app := &app.KohamApp{
		App: f_app,
		Routes: &routes.Routes{
			BaseValidator: &base_validators.BaseValidator{
				TokenService: &services.TokenService{},
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
			"POST",
			test.route,
			nil,
		)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("x-service-id", "d8c3eed5-8eda-441e-bcc1-16fab23b3ab7")

		// The -1 disables request latency.
		res, _ := koham_app.Test(req, -1)

		log.Println("Response", res)
		ctr := &dto.CreateTokenResponse{}
		err := json.NewDecoder(res.Body).Decode(ctr)
		if err != nil {
			log.Println(err)
		}
		log.Println("Response body", ctr)

		// Verify if the status code is as expected
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)
		// Verify, that the reponse body equals the expected body
		assert.Equalf(t, test.expectedBody, ctr, test.description)
	}
}
