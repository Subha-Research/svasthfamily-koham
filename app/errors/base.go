package errors

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

var DefaultErrorHandler = func(c *fiber.Ctx, err *fiber.Error) error {

	var error_json map[string]interface{}

	if _, ok := ErrorEnums[err.Message]; ok {
		error_json = map[string]interface{}{
			"status_code":   err.Code,
			"error_code":    err.Message,
			"error_message": ErrorEnums[err.Message].ErrorMessage,
		}
	} else {
		error_json = map[string]interface{}{
			"status_code":   500,
			"error_code":    "KSE-5001",
			"error_message": "Internal Server Error",
		}
		log.Printf("Error enums does not contain key %s", err.Message)
	}

	return c.Status(err.Code).JSON(error_json)
}
