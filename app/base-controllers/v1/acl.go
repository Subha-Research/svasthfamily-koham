package base_controllers

import (
	"github.com/Subha-Research/koham/app/interfaces"
	sf_controllers "github.com/Subha-Research/koham/app/svasthfamily/controllers/v1"
	"github.com/gofiber/fiber/v2"
)

type BaseACLController struct {
}

func (acl BaseACLController) Get(c *fiber.Ctx) error {
	stakeholder_type := c.Params("stakeholder_type")
	user_type := c.Params("user_type")

	// CLEANUPS:: Remove hardcoded values
	if stakeholder_type == "family" && user_type == "users" {
		return interfaces.IRequest.Get(sf_controllers.ACLController{}, c)
	} else {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request parameters")
	}
}

func (acl BaseACLController) Post(c *fiber.Ctx) error {
	stakeholder_type := c.Params("stakeholder_type")
	user_type := c.Params("user_type")

	// CLEANUPS:: Remove hardcoded values
	if stakeholder_type == "family" && user_type == "users" {
		return interfaces.IRequest.Post(sf_controllers.ACLController{}, c)
	} else {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request parameters")
	}
}

func (acl BaseACLController) Put(c *fiber.Ctx) error {
	return nil
}

func (acl BaseACLController) Delete(c *fiber.Ctx) error {
	return nil
}
