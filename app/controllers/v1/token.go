package controllers

import (
	"github.com/Subha-Research/svasthfamily-koham/app/errors"
	services "github.com/Subha-Research/svasthfamily-koham/app/services/v1"
	validators "github.com/Subha-Research/svasthfamily-koham/app/validators"
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

		tokenrb := new(validators.TokenRequestBody)
		if err := c.BodyParser(tokenrb); err != nil {
			// If any error in body parsing of fiber
			// So we return fiber error
			return errors.DefaultErrorHandler(c, fiber.NewError(400, "Body Parsing failed"))
		}
		// TODO:: Complete this method
	}
	ts := services.TokenService{}
	response, err := ts.CreateToken(f_user_id)
	if err != nil {
		return err
	}

	// TODO:: Update return to return json that will contain token key and expiry
	return c.Status(fiber.StatusCreated).JSON(response)
}

func (tc TokenController) Put(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).SendString("Token PUT called")
}

func (tc TokenController) Delete(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).SendString("Token DELETE called")
}
