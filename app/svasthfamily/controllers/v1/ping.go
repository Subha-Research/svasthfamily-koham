package sf_controllers

import (
	pariwar_services "github.com/Subha-Research/koham/app/svasthfamily/services/v1"
	sf_validators "github.com/Subha-Research/koham/app/svasthfamily/validators"
	"github.com/gofiber/fiber/v2"
)

type Person struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

func PingHandler(c *fiber.Ctx) error {
	p := new(Person)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	err := sf_validators.ValidatePing(p.Name)
	if err != nil {
		return err
	}

	pariwar_services.PingHandler(p.Name, p.Pass)
	return c.Status(fiber.StatusOK).SendString("Ping is working.")
}

func PingGetHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("Ping is working.")
}
