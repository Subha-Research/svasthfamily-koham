package base_controllers

import (
	"github.com/Subha-Research/koham/app/interfaces"
	sf_controllers "github.com/Subha-Research/koham/app/svasthfamily/controllers/v1"
	"github.com/gofiber/fiber/v2"
)

type BaseTokenController struct {
}

func (btc BaseTokenController) Get(c *fiber.Ctx) error {
	stakeholder_type := c.Params("stakeholder_type")
	user_type := c.Params("user_type")

	// CLEANUPS:: Remove hardcoded values
	if stakeholder_type == "family" && user_type == "users" {
		return interfaces.IRequest.Get(sf_controllers.TokenController{}, c)
	}
	return nil
}

func (btc BaseTokenController) Post(c *fiber.Ctx) error {
	stakeholder_type := c.Params("stakeholder_type")
	user_type := c.Params("user_type")

	// CLEANUPS:: Remove hardcoded values
	if stakeholder_type == "family" && user_type == "users" {
		return interfaces.IRequest.Post(sf_controllers.TokenController{}, c)
	}
	return nil
}

func (btc BaseTokenController) Put(c *fiber.Ctx) error {
	return nil
}

func (btc BaseTokenController) Delete(c *fiber.Ctx) error {
	return nil
}
