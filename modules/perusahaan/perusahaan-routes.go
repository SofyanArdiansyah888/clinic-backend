package perusahaan

import (
	"backend/config"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(api fiber.Router) {
	controller := NewPerusahaanController(NewPerusahaanService(NewPerusahaanRepository(config.DB)))
	pasienRoutes := api.Group("/perawatan")
	pasienRoutes.Post("/", controller.Store)
	pasienRoutes.Put("/:id", controller.Update)
	pasienRoutes.Get("/:id", controller.Show)
}
