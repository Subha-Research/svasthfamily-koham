package base_controllers

import (
	sf_controllers "github.com/Subha-Research/koham/app/svasthfamily/controllers/v1"
	"github.com/gofiber/fiber/v2"
)

type BaseACLController struct {
}

func (acl *BaseACLController) GetHandler(c *fiber.Ctx) error {
	stakeholder_type := c.Params("stakeholder_type")
	user_type := c.Params("user_type")

	if stakeholder_type == "family" && user_type == "users" {
		controller := sf_controllers.ACLController{}
		return controller.Get(c)
	} else {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request parameters")
	}
}
