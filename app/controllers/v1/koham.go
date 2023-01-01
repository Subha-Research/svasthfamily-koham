package controllers

import (
	"github.com/Subha-Research/svasthfamily-koham/app/interfaces"
	"github.com/gofiber/fiber/v2"
)

type BaseController struct {
	TokenController *TokenController
	ACLController   *ACLController
}

func (bc BaseController) GetHandler(c *fiber.Ctx) error {
	resource_type := c.Params("resource_type")

	// CLEANUPS:: Remove hardcoded values
	if resource_type == "tokens" {
		return interfaces.IRequest.Get(bc.TokenController, c)
	} else if resource_type == "acls" {
		return interfaces.IRequest.Get(bc.ACLController, c)
	} else {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request parameters")
	}
}

func (bc BaseController) PostHandler(c *fiber.Ctx) error {
	resource_type := c.Params("resource_type")

	// CLEANUPS:: Remove hardcoded values
	if resource_type == "tokens" {
		return interfaces.IRequest.Post(bc.TokenController, c)
	} else if resource_type == "acls" {
		return interfaces.IRequest.Post(bc.ACLController, c)
	} else {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request parameters")
	}
}

func (bc BaseController) PutHandler(c *fiber.Ctx) error {
	resource_type := c.Params("resource_type")

	// CLEANUPS:: Remove hardcoded values
	if resource_type == "tokens" {
		return interfaces.IRequest.Put(bc.TokenController, c)
	} else if resource_type == "acls" {
		return interfaces.IRequest.Put(bc.ACLController, c)
	} else {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request parameters")
	}
	//return nil
}

func (bc BaseController) DeleteHandler(c *fiber.Ctx) error {
	resource_type := c.Params("resource_type")

	// CLEANUPS:: Remove hardcoded values
	if resource_type == "tokens" {
		return interfaces.IRequest.Delete(bc.TokenController, c)
	} else if resource_type == "acls" {
		return interfaces.IRequest.Delete(bc.ACLController, c)
	} else {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request parameters")
	}
}
