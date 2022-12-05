package sf_controllers

import "github.com/gofiber/fiber/v2"

type TokenController struct {
}

func (tc *TokenController) Get(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("Get Family token.")
}

func (tc *TokenController) Post(c *fiber.Ctx) error {
	return nil
}
