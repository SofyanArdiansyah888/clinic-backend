package antrian

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(api fiber.Router) {
	controller := NewGenerateNumberController()
	group := api.Group("/generate-number")
	group.Get("/", controller.Show)
}
