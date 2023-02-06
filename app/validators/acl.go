package validators

import (
	"reflect"
	"strings"

	"github.com/Subha-Research/svasthfamily-koham/app/constants"
	"github.com/Subha-Research/svasthfamily-koham/app/errors"
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

type ChildUserAccess struct {
	ChildUserId string    `json:"family_child_user_id" validate:"required,uuid4_rfc4122"`
	AccessEnums []float64 `json:"access_enums" validate:"omitempty,min=1,dive,number"`
}

type ACLPostBody struct {
	AccessList     []ChildUserAccess `json:"access_list" validate:"required,min=1,dive"`
	ParentUserID   string            `json:"family_parent_user_id" validate:"required,uuid4_rfc4122"`
	FamilyID       string            `json:"family_id" validate:"omitempty,uuid_rfc4122"`
	FamilyMemberID string            `json:"family_member_id" validate:"omitempty,uuid_rfc4122"`
	IsParentHead   *bool             `json:"is_parent_head" validate:"required"`
}

type ACLPutBody struct {
	//In Future we may Required "ChildUserAccessList", "ParentUserID","RoleEnum" for implementing the Transfer Access feature.
	Access         *ChildUserAccess `json:"access" validate:"omitempty,required_with=ParentUserID"`
	ParentUserID   string           `json:"family_parent_user_id" validate:"omitempty,required_with=Access,uuid4_rfc4122"`
	FamilyID       string           `json:"family_id" validate:"omitempty,uuid_rfc4122"`
	FamilyMemberID string           `json:"family_member_id" validate:"omitempty,uuid_rfc4122"`
}

type ACLValidator struct {
}

func (av *ACLValidator) validateAccess(a_enums []float64) bool {
	// Run a loop to build freq_hash_map
	freq_hash_map := map[float64]int{}

	for k := range constants.CHILD_DEFAULT_ACCESS {
		freq_hash_map[k] = 1
	}

	for i := 0; i < len(a_enums); i++ {
		// Type assertion to int from interface
		if freq_hash_map[a_enums[i]] != 1 {
			return false
		}
	}
	return true
}

func (av *ACLValidator) ValidateACLPostBody(aclpb ACLPostBody, f_user_id string, is_head_acl_request bool) error {
	validate.RegisterTagNameFunc(ExtractTagName)
	err := validate.Struct(aclpb)
	error_data := map[string]string{
		"key": "role",
	}

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			error_data["key"] = err.Field()
			return errors.KohamError("SFKSE-4006", error_data)
		}
	}

	if !is_head_acl_request {
		if aclpb.FamilyID == "" {
			error_data["key"] = "family_id"
			return errors.KohamError("SFKSE-4006", error_data)
		} else if aclpb.FamilyMemberID == "" {
			error_data["key"] = "family_member_id"
			return errors.KohamError("SFKSE-4006", error_data)
		}
	}

	access_list := aclpb.AccessList
	for i := 0; i < len(access_list); i++ {
		access_enums := access_list[i].AccessEnums
		if access_enums != nil {
			is_all_access_present := av.validateAccess(access_enums)

			if !is_all_access_present {
				error_data["key"] = "access_enums"
				return errors.KohamError("SFKSE-4006", error_data)
			}
		}
	}
	return nil
}

func (av *ACLValidator) ValidateACLPutBody(aclputb ACLPutBody, f_user_id string) error {
	validate.RegisterTagNameFunc(ExtractTagName)
	err := validate.Struct(aclputb)

	error_data := map[string]string{
		"key": "family_child_user_id",
	}

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			error_data["key"] = err.Field()
			return errors.KohamError("SFKSE-4006", error_data)
		}
	}
	if aclputb.Access != nil {
		access := aclputb.Access
		access_enums := access.AccessEnums
		if access_enums == nil {
			error_data["key"] = "access_enums"
			return errors.KohamError("SFKSE-4006", error_data)
		}
		is_all_access_present := av.validateAccess(access_enums)

		if !is_all_access_present {
			error_data["key"] = "access_enums"
			return errors.KohamError("SFKSE-4006", error_data)
		}
	}
	return nil
}
