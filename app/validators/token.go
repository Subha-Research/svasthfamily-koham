package validators

import (
	"github.com/Subha-Research/svasthfamily-koham/app/enums"
)

type TokenPostBody struct {
	IsHead bool `json:"is_head" validate:"required"`
}

type TokenValidator struct {
}
type TokenRequestBody struct {
	Childmember_id string  `json:"parent_member_id" validate:"required,uuid4_rfc4122"`
	AccessEnum     float64 `json:"role" validate:"required,number"`
}

func (tv *TokenValidator) ValidateTokenRequestbody(child_member_id string, acessnumberreceived float64) error {
	d := enums.Accesses

	for k := range d {

		if k == acessnumberreceived {

		}
	}

	// var validate = validator.New()
	// err := validate.Struct(tpb)

	// if err != nil {
	// 	for _, err := range err.(validator.ValidationErrors) {
	// 		// var element errorResponse
	// 		// // fmt.Printf("*** %#v \n", err)
	// 		// element.MissingField = err.StructNamespace()
	// 		// element.Tag = err.ActualTag()
	// 		// element.Value = err.Param()
	// 		// element.Message = fmt.Sprintf("Invalid value of %s", err.StructNamespace())
	// 		// errors = append(errors, &element)
	// 		_ := fmt.Sprintf("%s", err.StructNamespace())
	// 		return fiber.NewError(fiber.StatusBadRequest, "KDE-4001")
	// 	}
	// }
	return nil
}
