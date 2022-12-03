package validators

import "github.com/gofiber/fiber/v2"

func ValidatePing(name string) error {
	if len(name) < 10 {
		return fiber.NewError(fiber.StatusBadRequest, "Length size is less than 10")
	}
	return nil
}
