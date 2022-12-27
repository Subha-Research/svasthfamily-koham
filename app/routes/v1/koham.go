package routes

import (
	base_validators "github.com/Subha-Research/svasthfamily-koham/app/base-validators"
	controllers "github.com/Subha-Research/svasthfamily-koham/app/controllers/v1"
	"github.com/Subha-Research/svasthfamily-koham/app/errors"
	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	Controller controllers.BaseController
	Validator  base_validators.BaseValidator
}

func (r *Routes) SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	v1 := api.Group("/family/users/:user_id/:resource_type")

	// Fiber middleware to validate headers.
	v1.Use("/", func(c *fiber.Ctx) error {
		// Validate headers if headers has required keys or not.
		err := r.Validator.ValidateHeaders(c)
		if err != nil {
			return errors.DefaultErrorHandler(c, err)
		}
		return c.Next()
	})
	v1.Get("/", r.Controller.GetHandler)
	v1.Post("/", r.Controller.PostHandler)
}
