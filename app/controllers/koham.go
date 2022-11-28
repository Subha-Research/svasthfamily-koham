package controllers

import (
	"github.com/Subha-Research/pariwar-koham/app/services/v1"
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

	services.PingHandler(p.Name, p.Pass)
	return c.Status(fiber.StatusOK).SendString("Ping is working.")
}
