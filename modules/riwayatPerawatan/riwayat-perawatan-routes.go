package riwayatPerawatan

import (
	"backend/config"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(api fiber.Router) {
	controller := NewRiwayatPerawatanController(NewRiwayatPerawatanService(NewRiwayatPerawatanRepository(config.DB)))
	pasienRoutes := api.Group("/pasien")
	pasienRoutes.Get("/", controller.Index)
	pasienRoutes.Get("/:id", controller.Show)
	pasienRoutes.Post("/", controller.Store)
	pasienRoutes.Put("/:id", controller.Update)
	pasienRoutes.Delete("/:id", controller.Delete)
}
