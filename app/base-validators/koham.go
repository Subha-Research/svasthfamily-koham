package base_validators

import (
	"log"
	"strings"

	"github.com/Subha-Research/svasthfamily-koham/app/errors"
	"github.com/Subha-Research/svasthfamily-koham/app/services/v1"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type BaseValidator struct {
}

const XServiceID = "d8c3eed5-8eda-441e-bcc1-16fab23b3ab7"

func (bv *BaseValidator) ValidateHeaders(c *fiber.Ctx) error {
	if c.Get("Content-Type") != "application/json" {
		return errors.KohamError("KSE-4001")
	}
	is_validate := c.Params("validate")
	resource_type := c.Params("resource_type")
	user_id := c.Params("user_id")

	if resource_type == "tokens" && is_validate == "" {
		log.Println(c.Get("x-service-id"))
		x_service_id, err := uuid.Parse(c.Get("x-service-id"))
		if err != nil {
			return errors.KohamError("KSE-4002")
		}
		// Validate if the x-service-token is supported by
		// CLEANUPS:: remove hardcoded x_service_id
		if x_service_id.String() != XServiceID {
			return errors.KohamError("KSE-4005")
		}
	} else if resource_type == "acls" || (resource_type == "tokens" && is_validate == "validate") {
		auth := c.Get("Authorization")
		if auth == "" {
			return errors.KohamError("KSE-4003")
		}
		token := strings.Split(auth, "Bearer ")
		if token[1] == "" {
			return errors.KohamError("KSE-4004")
		} else {
			// TRY to decode JWT token
			// return nil
			ts := services.TokenService{}
			err := ts.ParseToken(token[1], user_id)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
