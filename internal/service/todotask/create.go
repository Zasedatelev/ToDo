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
			"message": "Invalid input",
		})
		return
	}

	c.Status(fiber.StatusCreated)
	s.repo.Create(reqCtx, task)

}
