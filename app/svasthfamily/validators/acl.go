package sf_validators

import (
	"reflect"
	"strings"

	"github.com/Subha-Research/koham/app/errors"
	validator "github.com/go-playground/validator/v10"
)

var validate = validator.New()
var ExtractTagName = func(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	if name == "-" {
		return ""
	}

	return name
}

type ACLPostBody struct {
	ChildMemberIDS []string `json:"child_member_ids" validate:"required,min=1,dive,uuid4_rfc4122"`
	ParentMemberID string   `json:"parent_member_id" validate:"required,uuid4_rfc4122"`
	AccessEnums    []int    `json:"accesses" validate:"required,min=1,dive,number"`
}

type ACLPutBody struct {
	ChildMemberIDS []string `json:"child_member_ids" validate:"required,min=1,dive,uuid4_rfc4122"`
	ParentMemberID string   `json:"parent_member_id" validate:"required,uuid4_rfc4122"`
	AccessEnums    []int    `json:"accesses" validate:"required,min=1,dive,number"`
}

type ACLValidator struct {
}

func (av *ACLValidator) ValidateACLPostBody(aclpb ACLPostBody) error {
	validate.RegisterTagNameFunc(ExtractTagName)
	err := validate.Struct(aclpb)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			error_data := map[string]string{
				"key": err.Field(),
			}
			return errors.KohamError("KSE-4006", error_data)
		}
	}
	return nil
}
