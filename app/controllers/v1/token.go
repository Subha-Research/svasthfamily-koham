package controllers

import (
	"github.com/Subha-Research/svasthfamily-koham/app/errors"
	"github.com/Subha-Research/svasthfamily-koham/app/interfaces"
	validators "github.com/Subha-Research/svasthfamily-koham/app/validators"
	"github.com/gofiber/fiber/v2"
)

type TokenController struct {
	Validator *validators.TokenValidator
	IService  interfaces.ITokenService
}

func (tc TokenController) Get(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("Get Family token.")
}

func (tc TokenController) Post(c *fiber.Ctx) error {
	token := c.Locals("token").(*string)
	f_user_id := c.Params("user_id")
	opt_param := c.Params("validate")

	if opt_param == "validate" {
		tokenrb := new(validators.TokenRequestBody)
		err := c.BodyParser(tokenrb)
		if err != nil {
			// If any error in body parsing of fiber
			// So we return fiber error
			return errors.DefaultErrorHandler(c, fiber.NewError(400, "Body Parsing failed"))
		}
		// tv := validators.TokenValidator{}
		err_rb := tc.Validator.ValidateTokenRequestbody(*tokenrb)
		if err_rb != nil {
			// TODO:: Make koham error
			return err_rb
		}
		response, err := tc.IService.ValidateTokenAccess(token, f_user_id, *tokenrb)
		if err != nil {
			return errors.DefaultErrorHandler(c, err)
		}
		return c.Status(fiber.StatusOK).JSON(response)
	} else if opt_param == "" {
		response, err := tc.IService.CreateToken(f_user_id)
		if err != nil {
			return errors.DefaultErrorHandler(c, err)
		}
		return c.Status(fiber.StatusCreated).JSON(response)
	}
	return c.Status(fiber.StatusNotFound).SendString("404 URL not found")

}

func (tc TokenController) Put(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).SendString("Token PUT called")
}

func (tc TokenController) Delete(c *fiber.Ctx) error {
	// TODO:: Call service
	f_user_id := c.Params("user_id")
	token := c.Locals("token").(*string)
	err := tc.IService.DeleteToken(&f_user_id, token)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).SendString("Token DELETE called")
}
