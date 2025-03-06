package todotask

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Zasedatelev/ToDo.git/model"
	"github.com/gofiber/fiber"
)

func (s *serv) Update(c *fiber.Ctx) {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println("Conversion error:", err)
		return
	}
	task := model.Task{}
	reqCtx := c.Context()

	if err := c.BodyParser(&task); err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input"})
		return
	}

	err = s.repo.Update(reqCtx, task, int32(id))
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
		log.Panicf("Internal Server Error: %v\n", err)
	}
	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Task updated",
	})

}
