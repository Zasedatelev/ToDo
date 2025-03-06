package todotask

import (
	"github.com/gofiber/fiber"
)

func (s *serv) Get(c *fiber.Ctx) {

	reqCtx := c.Context()

	tasks, err := s.repo.Get(reqCtx)

	if err != nil {
		c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data not found",
		})
	}
	c.Set("Content-Type", "application/json")
	c.Status(fiber.StatusOK).JSON(tasks)

}
