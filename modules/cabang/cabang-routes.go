package cabang

import (
	"backend/config"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(api fiber.Router) {
	repository := NewCabangRepository(config.DB)
	service := NewCabangService(repository)
	controller := NewCabangController(service)

	group := api.Group("/cabang")
	group.Get("/", controller.Index)
	group.Get("/:id", controller.Show)
	group.Post("/", controller.Store)
	group.Put("/:id", controller.Update)
	group.Delete("/:id", controller.Delete)
}