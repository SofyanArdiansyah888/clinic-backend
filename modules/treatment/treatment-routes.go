package treatment

import (
	"backend/config"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(api fiber.Router) {
	repository := NewTreatmentRepository(config.DB)
	service := NewTreatmentService(repository)
	controller := NewTreatmentController(service)

	group := api.Group("/treatment")
	group.Get("/", controller.Index)
	group.Get("/:id", controller.Show)
	group.Post("/", controller.Store)
	group.Put("/:id", controller.Update)
	group.Delete("/:id", controller.Delete)

}
