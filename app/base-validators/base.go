package base_validators

import (
	"strings"

	"github.com/Subha-Research/svasthfamily-koham/app/constants"
	"github.com/Subha-Research/svasthfamily-koham/app/errors"
	"github.com/Subha-Research/svasthfamily-koham/app/interfaces"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type BaseValidator struct {
	ITokenService interfaces.ITokenService
}

const XServiceID = "d8c3eed5-8eda-441e-bcc1-16fab23b3ab7"

func (bv *BaseValidator) ValidateHeaders(c *fiber.Ctx) (*string, error) {
	if c.Get("Content-Type") != "application/json" {
		return nil, errors.KohamError("KSE-4001")
	}
	opt_param := c.Params("opt")
	req_method := c.Method()
	resource_type := c.Params("resource_type")
	user_id := c.Params("user_id")
	strategy, err := bv.headerValidationStrategy(resource_type, opt_param, req_method)
	if err != nil {
		return nil, err
	}

	if *strategy == constants.HEADER_VALIDATOR_STRATEGY["service"] {
		x_service_id, err := uuid.Parse(c.Get("x-service-id"))
		if err != nil {
			return nil, errors.KohamError("KSE-4002")
		}
		// Validate if the x-service-token is supported by
		// CLEANUPS:: remove hardcoded x_service_id
		if x_service_id.String() != XServiceID {
			return nil, errors.KohamError("KSE-4005")
		}
	} else if *strategy == constants.HEADER_VALIDATOR_STRATEGY["authorization"] {
		auth := c.Get("Authorization")
		if auth == "" {
			return nil, errors.KohamError("KSE-4003")
		}
		token := strings.Split(auth, "Bearer ")
		if token[1] == "" {
			return nil, errors.KohamError("KSE-4004")
		} else {
			err := bv.ITokenService.ParseToken(token[1], user_id)
			if err != nil {
				return nil, err
			}
		}
		return &token[1], nil
	}
	return nil, nil
}

func (bv *BaseValidator) headerValidationStrategy(resource_type string, opt string, req_method string) (*string, error) {
	var strategy string
	switch true {
	case resource_type == "tokens" && opt == "" && (req_method == "POST" || req_method == "GET"):
		strategy = constants.HEADER_VALIDATOR_STRATEGY["service"]
	case resource_type == "acls" && opt == "head":
		strategy = constants.HEADER_VALIDATOR_STRATEGY["service"]
	case resource_type == "tokens" && opt == "validate":
		strategy = constants.HEADER_VALIDATOR_STRATEGY["authorization"]
	case resource_type == "tokens" && req_method == "DELETE":
		strategy = constants.HEADER_VALIDATOR_STRATEGY["authorization"]
	case resource_type == "acls" && opt == "":
		strategy = constants.HEADER_VALIDATOR_STRATEGY["authorization"]
	default:
		return nil, errors.KohamError("KSE-4014")
	}
	return &strategy, nil
}
