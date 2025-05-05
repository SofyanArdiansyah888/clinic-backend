package antrian

import (
	"backend/config"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(api fiber.Router) {
	repository := NewAntrianRepository(config.DB)
	service := NewAntrianService(repository)
	controller := NewAntrianController(service)

	group := api.Group("/antrian")
	//group.Put("/:id/update-status", controller.UpdateStatus)
	group.Get("/", controller.Index)
	group.Get("/:id", controller.Show)
	group.Post("/", controller.Store)
	group.Put("/:id", controller.Update)
	group.Delete("/:id", controller.Delete)

}
