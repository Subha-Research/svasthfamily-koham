package sf_controllers

import (
	"log"

	sf_services "github.com/Subha-Research/koham/app/svasthfamily/services/v1"
	sf_validators "github.com/Subha-Research/koham/app/svasthfamily/validators"
	"github.com/gofiber/fiber/v2"
)

func PingHandler(c *fiber.Ctx) error {
	type Person struct {
		NameA string `json:"name_a" form:"name_a"`
		Pass  string `json:"pass" form:"pass"`
	}

	p := new(Person)

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	log.Println("HERE ", p.NameA)
	err := sf_validators.ValidatePing(p.NameA)
	if err != nil {
		return err
	}

	err_service := sf_services.PingHandler(p.NameA, p.Pass)
	if err_service != nil {
		return err_service
	}

	return c.Status(fiber.StatusOK).SendString("Ping POST is working.")
}

func PingGetHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("Ping is not working.")
}
