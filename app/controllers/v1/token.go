package controllers

import (
	services "github.com/Subha-Research/svasthfamily-koham/app/services/v1"
	"github.com/gofiber/fiber/v2"
)

type TokenController struct {
}

func (tc TokenController) Get(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("Get Family token.")
}

func (tc TokenController) Post(c *fiber.Ctx) error {
	f_user_id := c.Params("user_id")
	if c.Params("resource_type") == "tokens" && c.Params("resource_type") == "validate" {
		// TODO:: Complete this method
	}
	ts := services.TokenService{}
	ts.CreateToken(f_user_id)

	// tpb := new(validators.TokenPostBody)
	// if err := c.BodyParser(tpb); err != nil {
	// 	// If any error in body parsing
	// 	return errors.DefaultErrorHandler(c,
	// 		fiber.NewError(fiber.StatusBadRequest, "Body parsing failed."))
	// }

	// // Request body validation
	// tokenValidator := validators.TokenValidator{}
	// err := tokenValidator.ValidatePostBody(*tpb)
	// if err != nil {
	// 	return errors.DefaultErrorHandler(c, err)
	// }

	// TODO:: Update return to return json that will contain token key and expiry
	return c.Status(fiber.StatusCreated).SendString("Token creation successful")
}

func (tc TokenController) Put(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).SendString("Token PUT called")
}

func (tc TokenController) Delete(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).SendString("Token DELETE called")
}
