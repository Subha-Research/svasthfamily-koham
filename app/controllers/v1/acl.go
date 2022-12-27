package controllers

import (
	"github.com/Subha-Research/svasthfamily-koham/app/errors"
	services "github.com/Subha-Research/svasthfamily-koham/app/services/v1"
	validators "github.com/Subha-Research/svasthfamily-koham/app/validators"
	"github.com/gofiber/fiber/v2"
)

// type IACLController interface {
// 	Get(*fiber.Ctx) error
// 	Post(*fiber.Ctx) error
// 	Put(*fiber.Ctx) error
// 	Delete(*fiber.Ctx) error
// }

type ACLController struct {
	Validator validators.ACLValidator
	Service   services.ACLService
}

func (acl ACLController) Get(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("Get family ACL")
}

func (acl ACLController) Post(c *fiber.Ctx) error {
	// Validate request body
	// Call service method build model data
	// TODO:: Validate token access before inserting
	// into database
	// Implement dependeny injection
	// Insert into database
	// Implement DTO
	f_user_id := c.Params("user_id")

	aclpb := new(validators.ACLPostBody)
	if err := c.BodyParser(aclpb); err != nil {
		// If any error in body parsing of fiber
		// So we return fiber error
		return errors.DefaultErrorHandler(c, fiber.NewError(400, "Body Parsing failed"))
	}

	// Request body validation
	// CLEANUP:: Access using interface
	err := acl.Validator.ValidateACLPostBody(*aclpb, f_user_id)
	if err != nil {
		return errors.DefaultErrorHandler(c, err)
	}

	// Call service
	if err := acl.Service.CreateSFRelationship(f_user_id, *aclpb); err != nil {
		return errors.DefaultErrorHandler(c, err)
	}
	return c.Status(fiber.StatusOK).SendString("POST family ACL")
}

func (acl ACLController) Put(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("PUT family ACL")
}

func (acl ACLController) Delete(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("DELETE family ACL")
}
