package base_validators

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type BaseValidator struct {
	XServiceID string
}

const XServiceID = "d8c3eed5-8eda-441e-bcc1-16fab23b3ab7"

func (bv *BaseValidator) ValidateHeaders(c *fiber.Ctx) *fiber.Error {
	if c.Get("Content-Type") != "application/json" {
		// err := errors.ErrorEnums["KSE-0001"]
		return fiber.NewError(fiber.StatusBadRequest, "KSE-4001")
	}
	resource_type := c.Params("resource_type")
	if resource_type == "tokens" {

		x_service_id, err := uuid.Parse(c.Get("x-service-id"))
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "KSE-4002")
		}
		// Validate if the x-service-token is supported by
		// CLEANUPS:: remove hardcoded x_service_id
		if x_service_id.String() != bv.XServiceID {
			return fiber.NewError(fiber.StatusUnauthorized, "KSE-4005")
		}
	} else if resource_type == "acls" {
		auth := c.Get("Authorization")
		if auth == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "KSE-4003")
		}
		token := strings.Split(auth, "Bearer ")
		if token[1] == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "KSE-4004")
		} else {
			// TRY to decode JWT token
			return nil
		}
	}
	return nil
}
