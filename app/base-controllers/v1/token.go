package base_controllers

import (
	sf_controllers "github.com/Subha-Research/koham/app/svasthfamily/controllers/v1"
	"github.com/gofiber/fiber/v2"
)

type BaseTokenController struct {
}

func (btc *BaseTokenController) GetHandler(c *fiber.Ctx) error {
	stakeholder_type := c.Params("stakeholder_type")
	user_type := c.Params("user_type")

	// CLEANUPS:: Remove hardcoded values
	if stakeholder_type == "family" && user_type == "users" {
		controller := sf_controllers.TokenController{}
		return controller.Get(c)
	}
	return nil
}

func (btc *BaseTokenController) PostHandler(c *fiber.Ctx) error {
	stakeholder_type := c.Params("stakeholder_type")
	user_type := c.Params("user_type")

	// CLEANUPS:: Remove hardcoded values
	if stakeholder_type == "family" && user_type == "users" {
		// Validate headers
		controller := sf_controllers.TokenController{}
		return controller.Post(c)
	}
	return nil
}
