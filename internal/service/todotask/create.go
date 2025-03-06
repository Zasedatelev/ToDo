package todotask

import (
	"github.com/Zasedatelev/ToDo.git/model"
	"github.com/gofiber/fiber"
)

func (s *serv) Create(c *fiber.Ctx) {
	task := model.Task{}
	reqCtx := c.Context()
	if err := c.BodyParser(&task); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid json",
		})
	}
	s.repo.Create(reqCtx, task)
}
