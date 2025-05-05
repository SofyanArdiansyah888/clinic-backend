package antrian

import (
	"backend/config"
	"backend/utils"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
}

func NewGenerateNumberController() *Controller {
	return &Controller{}
}

func (h *Controller) Show(c *fiber.Ctx) error {
	idParam := c.Query("key")
	lastId := utils.GenerateID(config.DB, idParam, false)
	return c.JSON(fiber.Map{
		"last_id": lastId,
		"params":  idParam,
	})
}
