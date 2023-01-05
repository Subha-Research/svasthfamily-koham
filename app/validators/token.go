package validators

import (
	"github.com/Subha-Research/svasthfamily-koham/app/constants"
	"github.com/Subha-Research/svasthfamily-koham/app/errors"
	validator "github.com/go-playground/validator/v10"
)

// type TokenPostBody struct {
// 	IsHead bool `json:"is_head" validate:"required"`
// }

type TokenValidator struct {
}

type ValidateTokenRB struct {
	ChildUserID string  `json:"child_user_id" validate:"required,uuid4_rfc4122"`
	AccessEnum  float64 `json:"access_enum" validate:"required,number"`
}

func (tv *TokenValidator) ValidateTokenRequestbody(rb ValidateTokenRB) error {
	err := validate.Struct(rb)
	error_data := map[string]string{
		"key": "role",
	}

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			error_data["key"] = err.Field()
			return errors.KohamError("KSE-4006", error_data)
		}
	}

	// Validate access enums
	d := constants.HEAD_DEFAULT_ACCESS
	for k := range d {
		if k == rb.AccessEnum {
			return nil
		}

	}
	error_data["key"] = "access"
	return errors.KohamError("KSE-4006", error_data)
}
