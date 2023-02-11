package routes

import (
	base_validators "github.com/Subha-Research/svasthfamily-koham/app/base-validators"
	controllers "github.com/Subha-Research/svasthfamily-koham/app/controllers/v1"
	"github.com/Subha-Research/svasthfamily-koham/app/errors"
	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	BaseController *controllers.BaseController
	BaseValidator  *base_validators.BaseValidator
}

func (r *Routes) SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	v1 := api.Group("/users/:user_id/:resource_type/:opt?")

	// Fiber middleware to validate headers.
	v1.Use("/", func(c *fiber.Ctx) error {
		// Validate headers if headers has required keys or not.
		token, err := r.BaseValidator.ValidateHeaders(c)
		if err != nil {
			return errors.DefaultErrorHandler(c, err)
		}
		c.Locals("token", token)
		return c.Next()
	})
	v1.Get("/", r.BaseController.GetHandler)
	v1.Post("/", r.BaseController.PostHandler)
	v1.Put("/", r.BaseController.PutHandler)
	v1.Delete("/", r.BaseController.DeleteHandler)
}
