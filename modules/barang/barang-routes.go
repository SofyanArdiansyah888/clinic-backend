package barang

import (
	"backend/config"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(api fiber.Router) {
	controller := NewController(NewBarangService(NewBarangRepository(config.DB)))
	barangRoutes := api.Group("/barang")
	barangRoutes.Get("/", controller.Index)
	barangRoutes.Get("/:id", controller.Show)
	barangRoutes.Post("/", controller.Store)
	barangRoutes.Put("/:id", controller.Update)
	barangRoutes.Delete("/:id", controller.Delete)
}
