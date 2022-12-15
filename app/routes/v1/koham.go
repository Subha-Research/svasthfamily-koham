package routes

import (
	base_validators "github.com/Subha-Research/koham/app/base-validators"
	sf_controllers "github.com/Subha-Research/koham/app/controllers/v1"
	"github.com/Subha-Research/koham/app/errors"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	bc := sf_controllers.BaseController{}

	v1 := api.Group("/family/users/:user_id/:resource_type")

	// Fiber middleware to validate headers.
	v1.Use("/", func(c *fiber.Ctx) error {
		// Validate headers if headers has required keys or not.
		bv := base_validators.BaseValidator{}
		err := bv.ValidateHeaders(c)
		if err != nil {
			return errors.DefaultErrorHandler(c, err)
		}
		return c.Next()
	})
	v1.Get("/", bc.GetHandler)
	v1.Post("/", bc.PostHandler)
}
