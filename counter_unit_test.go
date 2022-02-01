package main

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCounter(t *testing.T) {
	// Define a structure for specifying input and output data
	// of a single test case
	testCount := 0
	tests := []struct {
		counter      int    // our counter data
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		// Display test case
		{
			counter:      testCount,
			description:  "get HTTP status 200",
			route:        "/display",
			expectedCode: 200,
		},
		// Increase test case
		{
			counter:      testCount,
			description:  "get HTTP status 200",
			route:        "/increase",
			expectedCode: 200,
		},
		// Decrease test case
		{
			counter:      testCount,
			description:  "get HTTP status 200",
			route:        "/decrease",
			expectedCode: 200,
		},
	}

	// Define Fiber app.
	app := fiber.New()

	// Create route with GET method for test
	app.Get("/display", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprint(testCount))
	})
	app.Get("/increase", func(c *fiber.Ctx) error {
		testCount++
		return c.SendString(fmt.Sprint(testCount))
	})
	app.Get("/decrease", func(c *fiber.Ctx) error {
		testCount--
		return c.SendString(fmt.Sprint(testCount))
	})

	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route from the test case
		req := httptest.NewRequest("GET", test.route, nil)

		// Perform the request plain with the app,
		// the second argument is a request latency
		// (set to -1 for no latency)
		resp, _ := app.Test(req, 1)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
