package errors

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Error struct {
	StatusCode int    `json:"status_code"`
	ErrorCode  string `json:"error_code"`
	Message    string `json:"message"`
}

// Error makes it compatible with the `error` interface.
func (e *Error) Error() string {
	return e.Message
}

func ExtractDynamicVars(message string) []string {
	var result = []string{}
	words := strings.Split(message, " ")
	for i := 0; i < len(words); i++ {
		word := words[i]
		fc := word[0:1]
		lc := word[len(word)-1:]
		if fc == ":" && lc == ":" {
			result = append(result, word)
		}
	}
	return result
}

func ReplaceDynamicVars(error_msg string, data map[string]string) string {
	dynamic_vars := ExtractDynamicVars(error_msg)

	for i := 0; i < len(dynamic_vars); i++ {
		d_var := dynamic_vars[i]
		if _, ok := data[d_var]; ok {
			error_msg = strings.Replace(error_msg, d_var, data[d_var], 1)
		} else {
			return ""
		}
	}
	return error_msg
}

// NewError creates a new Error instance with an optional data
func NewError(error_code string, data ...map[string]string) *Error {
	err := &Error{
		ErrorCode:  error_code,
		StatusCode: 500,
		Message:    "Internal Server Error",
	}

	if _, ok := ErrorEnums[error_code]; ok {
		error_msg := ErrorEnums[error_code].ErrorMessage
		if (len(data)) > 0 {
			error_msg = ReplaceDynamicVars(error_msg, data[0])
		}

		if error_msg != "" {
			err := &Error{
				ErrorCode:  error_code,
				StatusCode: ErrorEnums[error_code].StatusCode,
				Message:    error_msg,
			}
			return err
		} else {
			return err
		}
	}
	return err
}

var DefaultErrorHandler = func(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *Error
	if errors.As(err, &e) {
		code = e.StatusCode
	}

	// Return status code with error json
	return c.Status(code).JSON(err)
}
