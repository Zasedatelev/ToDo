package todotask

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber"
)

func (s *serv) Delete(c *fiber.Ctx) {
	reqCtx := c.Context()

	id := c.Params("id")

	i, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Conversion error:", err)
		return
	}

	if err := s.repo.Delete(reqCtx, int32(i)); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
		log.Panicf("Internal Server Error: %v\n", err)
	}
	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Task deleted",
	})

}
