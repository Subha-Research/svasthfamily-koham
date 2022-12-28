package routes

import (
	base_validators "github.com/Subha-Research/svasthfamily-koham/app/base-validators"
	controllers "github.com/Subha-Research/svasthfamily-koham/app/controllers/v1"
	"github.com/Subha-Research/svasthfamily-koham/app/errors"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	bc := controllers.BaseController{}

	v1 := api.Group("/family/users/:user_id/:resource_type/:validate?")

	// Fiber middleware to validate headers.
	v1.Use("/", func(c *fiber.Ctx) error {
		// Validate headers if headers has required keys or not.
		bv := base_validators.BaseValidator{}
		token, err := bv.ValidateHeaders(c)

		if err != nil {
			return errors.DefaultErrorHandler(c, err)
		}
		c.Locals("token", token)
		return c.Next()
	})
	v1.Get("/", bc.GetHandler)
	v1.Post("/", bc.PostHandler)
	v1.Delete("/", bc.DeleteHandler)
}
