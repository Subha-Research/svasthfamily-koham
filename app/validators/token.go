package validators

import (
	"github.com/gofiber/fiber/v2"
)

type TokenPostBody struct {
	IsHead bool `json:"is_head" validate:"required"`
}

type TokenValidator struct {
}

func (tv *TokenValidator) ValidatePostBody(tpb TokenPostBody) *fiber.Error {
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
