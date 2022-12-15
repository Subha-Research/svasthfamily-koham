package tests

import (
	"io"
	"net/http"
	"testing"

	app "github.com/Subha-Research/koham/app"
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

	app := app.KohamApp{}
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
