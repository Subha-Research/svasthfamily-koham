package sf_controllers

import (
	"github.com/gofiber/fiber/v2"
)

type TokenController struct {
}

func (tc *TokenController) Get(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("Get Family token.")
}

func (tc *TokenController) Post(c *fiber.Ctx) error {
	// f_user_id := c.Params("user_id")

	// tpb := new(sf_validators.TokenPostBody)
	// if err := c.BodyParser(tpb); err != nil {
	// 	// If any error in body parsing
	// 	return errors.DefaultErrorHandler(c,
	// 		fiber.NewError(fiber.StatusBadRequest, "Body parsing failed."))
	// }

	// // Request body validation
	// tokenValidator := sf_validators.TokenValidator{}
	// err := tokenValidator.ValidatePostBody(*tpb)
	// if err != nil {
	// 	return errors.DefaultErrorHandler(c, err)
	// }

	return c.Status(fiber.StatusCreated).SendString("Token creation successful")
}
