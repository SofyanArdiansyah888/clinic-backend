package perawatan

import (
	"backend/config"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(api fiber.Router) {
	controller := NewPerawatanController(NewPerawatanService(NewPerawatanRepository(config.DB)))
	pasienRoutes := api.Group("/perawatan")
	pasienRoutes.Get("/", controller.Index)
	pasienRoutes.Get("/:id", controller.Show)
	pasienRoutes.Post("/", controller.Store)
	pasienRoutes.Put("/:id", controller.Update)
	pasienRoutes.Delete("/:id", controller.Delete)
}
