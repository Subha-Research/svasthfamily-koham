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

type ACLPostBody struct {
	ChildMemberAccessList []map[string]interface{} `json:"child_member_access_list" validate:"required,min=1"`
	ParentMemberID        string                   `json:"parent_member_id" validate:"required,uuid4_rfc4122"`
	RoleEnum              int                      `json:"role" validate:"required,number"`
}

type ACLPutBody struct {
	ChildMemberIDS []string `json:"child_member_ids" validate:"required,min=1,dive,uuid4_rfc4122"`
	ParentMemberID string   `json:"parent_member_id" validate:"required,uuid4_rfc4122"`
	AccessEnums    []int    `json:"accesses" validate:"required,min=1,dive,number"`
	RoleEnum       int      `json:"role" validate:"required,number"`
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

func (av *ACLValidator) validateAccess(a_enums []interface{}) bool {
	// Run a loop to build freq_hash_map
	freq_hash_map := map[float64]int{}

	for k := range enums.Accesses {
		freq_hash_map[k] = 1
	}

	for i := 0; i < len(a_enums); i++ {
		// Type assertion to int from interface
		if freq_hash_map[a_enums[i].(float64)] != 1 {
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

	access_list := aclpb.ChildMemberAccessList
	for i := 0; i < len(access_list); i++ {
		child_member_id := access_list[i]["child_member_id"].(string)
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
		access_enums := access_list[i]["access_enums"].([]interface{})
		is_all_access_present := av.validateAccess(access_enums)

		if !is_all_access_present {
			error_data["key"] = "access_enums"
			return errors.KohamError("KSE-4006", error_data)
		}
	}
	return nil
}
