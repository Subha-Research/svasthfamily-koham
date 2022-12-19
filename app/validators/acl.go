package validators

import (
	"reflect"
	"strings"

	sf_enums "github.com/Subha-Research/svasthfamily-koham/app/enums"
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

type ACLPostBody struct {
	ChildMemberAccessList []map[string]interface{} `json:"child_member_access_list" validate:"gt=0,dive,keys,eq=2,endkeys,required"`
	ParentMemberID        string                   `json:"parent_member_id" validate:"required,uuid4_rfc4122"`
	//AccessEnums    []int    `json:"accesses" validate:"required,min=1,dive,number"`
	RoleEnum int `json:"role" validate:"required,number"`
}

type ACLPutBody struct {
	ChildMemberIDS []string `json:"child_member_ids" validate:"required,min=1,dive,uuid4_rfc4122"`
	ParentMemberID string   `json:"parent_member_id" validate:"required,uuid4_rfc4122"`
	AccessEnums    []int    `json:"accesses" validate:"required,min=1,dive,number"`
	RoleEnum       int      `json:"role" validate:"required,number"`
}

type ACLValidator struct {
}

func (av *ACLValidator) vaidateRole(r_enum int) bool {
	for k := range sf_enums.Roles {
		if r_enum == k {
			return true
		}
	}
	return false
}

// func (av *ACLValidator) validateAccess(a_enums []int) bool {
// 	// Run a loop to build freq_hash_map
// 	freq_hash_map := map[int]int{}

// 	for k := range sf_enums.Accesses {
// 		freq_hash_map[k] = 1
// 	}

// 	for i := 0; i < len(a_enums); i++ {
// 		if freq_hash_map[a_enums[i]] != 1 {
// 			return false
// 		}
// 	}
// 	return true
// }

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
	// Validate if we support the given role and access received in request
	// LATER CLEANUP:: validate from service by fetching from cache / database
	// role := aclpb.RoleEnum
	// access_enums := aclpb.AccessEnums
	// is_all_access_present := av.validateAccess(access_enums)
	// if !is_all_access_present {
	// 	error_data := map[string]string{
	// 		"key": "accesses",
	// 	}
	// 	return errors.KohamError("KSE-4006", error_data)
	// }

	role_enum := aclpb.RoleEnum
	is_role_valid := av.vaidateRole(role_enum)
	if !is_role_valid {
		error_data := map[string]string{
			"key": "role",
		}
		return errors.KohamError("KSE-4006", error_data)
	}
	return nil
}
