package sf_controllers

import (
	"github.com/Subha-Research/koham/app/interfaces"
	"github.com/gofiber/fiber/v2"
)

type BaseController struct {
}

func (bc BaseController) GetHandler(c *fiber.Ctx) error {
	resource_type := c.Params("resource_type")

	// CLEANUPS:: Remove hardcoded values
	if resource_type == "tokens" {
		return interfaces.IRequest.Get(TokenController{}, c)
	} else if resource_type == "acls" {
		return interfaces.IRequest.Get(ACLController{}, c)
	} else {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request parameters")
	}
}

func (bc BaseController) PostHandler(c *fiber.Ctx) error {
	resource_type := c.Params("resource_type")

	// CLEANUPS:: Remove hardcoded values
	if resource_type == "tokens" {
		return interfaces.IRequest.Post(TokenController{}, c)
	} else if resource_type == "acls" {
		return interfaces.IRequest.Post(ACLController{}, c)
	} else {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request parameters")
	}
}

func (bc BaseController) PutHandler(c *fiber.Ctx) error {
	return nil
}

func (bc BaseController) DeleteHandler(c *fiber.Ctx) error {
	return nil
}
