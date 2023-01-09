package controllers

import (
	"github.com/Subha-Research/svasthfamily-koham/app/errors"
	"github.com/Subha-Research/svasthfamily-koham/app/interfaces"
	services "github.com/Subha-Research/svasthfamily-koham/app/services/v1"
	validators "github.com/Subha-Research/svasthfamily-koham/app/validators"
	"github.com/gofiber/fiber/v2"
)

type ACLController struct {
	Validator     *validators.ACLValidator
	Service       *services.ACLService
	ITokenService interfaces.ITokenService
}

func (acl ACLController) Get(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("Get family ACL")
}

func (acl ACLController) Post(c *fiber.Ctx) error {
	// TODO:: Validate token access before inserting
	// into database
	// Implement DTO for response
	token := c.Locals("token").(*string)
	f_user_id := c.Params("user_id")
	if token != nil {
		// Validate token access
		rb := validators.ValidateTokenRB{
			ChildUserID: f_user_id,
			AccessEnum:  101, // 101 means ADD_SFM access
		}
		_, err := acl.ITokenService.ValidateTokenAccess(token, f_user_id, rb)
		if err != nil {
			return errors.DefaultErrorHandler(c, err)
		}
	}

	aclpb := new(validators.ACLPostBody)
	if err := c.BodyParser(aclpb); err != nil {
		// If any error in body parsing of fiber
		// So we return fiber error
		return errors.DefaultErrorHandler(c, fiber.NewError(400, "Body Parsing failed"))
	}

	err := acl.Validator.ValidateACLPostBody(*aclpb, f_user_id)
	if err != nil {
		return errors.DefaultErrorHandler(c, err)
	}

	res, err := acl.Service.CreateAccessRelationship(f_user_id, token, *aclpb)

	if err != nil {
		return errors.DefaultErrorHandler(c, err)
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

func (acl ACLController) Put(c *fiber.Ctx) error {
	token := c.Locals("token").(*string)
	f_user_id := c.Params("user_id")
	if token != nil {
		// Validate token access
		rb := validators.ValidateTokenRB{
			ChildUserID: f_user_id,
			AccessEnum:  103, // 101 means ADD_SFM access
		}
		_, err := acl.ITokenService.ValidateTokenAccess(token, f_user_id, rb)
		if err != nil {
			return errors.DefaultErrorHandler(c, err)
		}
	}
	aclputb := new(validators.ACLPutBody)
	if err := c.BodyParser(aclputb); err != nil {
		// If any error in body parsing of fiber
		// So we return fiber error
		return errors.DefaultErrorHandler(c, fiber.NewError(400, "Body Parsing failed"))
	}
	acl_validator := validators.ACLValidator{}
	err := acl_validator.ValidateACLPutBody(*aclputb, f_user_id)
	if err != nil {
		return errors.DefaultErrorHandler(c, err)
	}
	update_doc_response, err := acl.Service.UpdateAccessRelationship(f_user_id, *aclputb)
	if err != nil {
		return errors.DefaultErrorHandler(c, err)
	}
	return c.Status(fiber.StatusOK).JSON(update_doc_response)
}

func (acl ACLController) Delete(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("DELETE family ACL")
}
