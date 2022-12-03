package controllers

import (
	"github.com/Subha-Research/koham/app/services/v1"
	"github.com/Subha-Research/koham/app/validators"
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

	err := validators.ValidatePing(p.Name)
	if err != nil {
		return err
	}

	services.PingHandler(p.Name, p.Pass)
	return c.Status(fiber.StatusOK).SendString("Ping is working.")
}

func PingGetHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("Ping is working.")
}
