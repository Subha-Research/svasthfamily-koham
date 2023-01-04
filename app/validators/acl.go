package validators

import (
	"reflect"
	"strings"

	enums "github.com/Subha-Research/svasthfamily-koham/app/enums"
	"github.com/Subha-Research/svasthfamily-koham/app/errors"
	validator "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var validate = validator.New()
var ExtractTagName = func(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	if name == "-" {
		return ""
	}
	return name
}

type ChildMemberAccess struct {
	ChildMemberId string    `json:"child_member_id" validate:"required,uuid4_rfc4122"`
	AccessEnums   []float64 `json:"access_enums" validate:"omitempty,min=1,dive,number"`
}

type ACLPostBody struct {
	AccessList     []ChildMemberAccess `json:"access_list" validate:"required,min=1,dive"`
	ParentMemberID string              `json:"parent_member_id" validate:"required,uuid4_rfc4122"`
	RoleEnum       int                 `json:"role" validate:"required,number"`
	IsParentHead   *bool               `json:"is_parent_head" validate:"required"`
}

type ACLPutBody struct {
	//In Future we may Required "ChildMemberAccessList", "ParentMemberID","RoleEnum" for implementing the Transfer Access feature.
	Access         ChildMemberAccess `json:"access" validate:"required"`
	ParentMemberID string            `json:"parent_member_id" validate:"required,uuid4_rfc4122"`
}

type ACLValidator struct {
}

func (av *ACLValidator) validateRole(r_enum int) (bool, string) {
	for k, v := range enums.Roles {
		if r_enum == k {
			return true, v
		}
	}
	return false, ""
}

func (av *ACLValidator) validateAccess(a_enums []float64) bool {
	// Run a loop to build freq_hash_map
	freq_hash_map := map[float64]int{}

	for k := range enums.CHILD_DEFAULT_ACCESS {
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

func (av *ACLValidator) ValidateACLPostBody(aclpb ACLPostBody, f_user_id string) error {
	validate.RegisterTagNameFunc(ExtractTagName)
	err := validate.Struct(aclpb)
	error_data := map[string]string{
		"key": "role",
	}

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			error_data["key"] = err.Field()
			return errors.KohamError("KSE-4006", error_data)
		}
	}

	role_enum := aclpb.RoleEnum
	is_role_valid, role_key := av.validateRole(role_enum)
	if !is_role_valid {
		error_data["key"] = "role"
		return errors.KohamError("KSE-4006", error_data)
	}

	// Validate if we support the given role and access received in request
	// LATER CLEANUP:: validate from service by fetching from cache / database
	// Validate all child_member_id and access_enums
	// Check if child_member_id is in uuid format or not
	// and access enums are in supported list

	access_list := aclpb.AccessList
	for i := 0; i < len(access_list); i++ {
		child_member_id := access_list[i].ChildMemberId
		c_id, err := uuid.Parse(child_member_id)
		if err != nil {
			error_data["key"] = "child_member_id"
			return errors.KohamError("KSE-4006", error_data)
		}

		if role_key != "FAMILY_HEAD" {
			if c_id.String() == f_user_id {
				return errors.KohamError("KSE-4008")
			}
		}
		access_enums := access_list[i].AccessEnums
		if access_enums != nil {
			is_all_access_present := av.validateAccess(access_enums)

			if !is_all_access_present {
				error_data["key"] = "access_enums"
				return errors.KohamError("KSE-4006", error_data)
			}
		}
	}
	return nil
}

func (av *ACLValidator) ValidateACLPutBody(aclputb ACLPutBody, f_user_id string) error {
	validate.RegisterTagNameFunc(ExtractTagName)
	err := validate.Struct(aclputb)

	error_data := map[string]string{
		"key": "child_member_id",
	}

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			error_data["key"] = err.Field()
			return errors.KohamError("KSE-4006", error_data)
		}
	}
	access := aclputb.Access
	child_member_id := access.ChildMemberId
	c_id, err := uuid.Parse(child_member_id)
	if err != nil {
		error_data["key"] = "child_member_id"
		return errors.KohamError("KSE-4006", error_data)
	}
	p_id, err := uuid.Parse(aclputb.ParentMemberID)
	if err != nil {
		error_data["key"] = "parent_member_id"
		return errors.KohamError("KSE-4006", error_data)
	}

	if c_id.String() == f_user_id || p_id.String() == f_user_id {
		return errors.KohamError("KSE-4008")
	}

	access_enums := access.AccessEnums
	is_all_access_present := av.validateAccess(access_enums)

	if !is_all_access_present {
		error_data["key"] = "access_enums"
		return errors.KohamError("KSE-4006", error_data)
	}

	return nil
}
