package sf_controllers

import "github.com/gofiber/fiber/v2"

type ACLController struct {
}

func (acl *ACLController) Get(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("Get family ACL")
}
