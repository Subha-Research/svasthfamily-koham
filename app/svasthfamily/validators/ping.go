package sf_validators

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func ValidatePing(name string) error {
	log.Println(name)
	if len(name) < 10 {
		return fiber.NewError(fiber.StatusBadRequest, "Length size is less than 10")
	}
	return nil
}
