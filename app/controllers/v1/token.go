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
	token := c.Locals("token").(*string)
	f_user_id := c.Params("user_id")
	ts := services.TokenService{}
	if c.Params("validate") == "validate" {

		tokenrb := new(validators.TokenRequestBody)
		err := c.BodyParser(tokenrb)
		if err != nil {
			// If any error in body parsing of fiber
			// So we return fiber error
			return errors.DefaultErrorHandler(c, fiber.NewError(400, "Body Parsing failed"))
		}
		tv := validators.TokenValidator{}
		err_rb := tv.ValidateTokenRequestbody(*tokenrb)
		if err_rb != nil {
			// TODO:: Make koham error
			return err_rb
		}
		// TODO:: Complete this method
		response, err := ts.ValidateTokenAccess(token, f_user_id, *tokenrb)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(response)
	}

	response, err := ts.CreateToken(f_user_id)
	if err != nil {
		return errors.DefaultErrorHandler(c, err)
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
