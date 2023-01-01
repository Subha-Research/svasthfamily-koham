package tests

import (
	"io"
	"net/http"
	"testing"

	app "github.com/Subha-Research/svasthfamily-koham/app"
	"github.com/stretchr/testify/assert"
)

type ACLRouteTest struct {
}

func (acltest *ACLRouteTest) TestCreateACL(t *testing.T) {
	tests := []struct {
		description string
		route       string
		// Expected output
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "ACL API testcase",
			route:         "/api/v1/family/users/8204a616-2131-4a64-97d0-ae3f2b9211be/acls",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "Ping is working.",
		},
	}

	app := &app.KohamApp{
		App: app.InitFiberApplication(),
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
