package service

import (
	"github.com/gofiber/fiber"
)

type ToDoService interface {
	Get(c *fiber.Ctx)
	Create(c *fiber.Ctx)
}
