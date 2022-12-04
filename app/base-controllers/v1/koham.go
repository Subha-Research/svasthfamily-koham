package base_controllers

import (
	"github.com/gofiber/fiber/v2"
)

type BaseController struct {
}

func (btc *BaseController) GetHandler(c *fiber.Ctx) error {
	resource_type := c.Params("resource_type")
	if resource_type == "tokens" {
		btc := BaseTokenController{}
		return btc.GetHandler(c)
	} else if resource_type == "acls" {
		acl_controller := BaseACLController{}
		return acl_controller.GetHandler(c)
	} else {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request parameters")
	}
}
