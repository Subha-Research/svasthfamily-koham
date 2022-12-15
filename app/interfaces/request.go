package interfaces

import "github.com/gofiber/fiber/v2"

type IRequest interface {
	Get(c *fiber.Ctx) error
	Post(c *fiber.Ctx) error
	Put(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
