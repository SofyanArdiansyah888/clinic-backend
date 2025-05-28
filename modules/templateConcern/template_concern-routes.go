package templateConcern

import (
	"backend/config"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(api fiber.Router) {
	repository := NewTemplateConcernRepository(config.DB)
	service := NewTemplateConcernService(repository)
	controller := NewTemplateConcernController(service)

	group := api.Group("/template-concern")
	group.Get("/", controller.Index)
	group.Get("/:id", controller.Show)
	group.Post("/", controller.Store)
	group.Put("/:id", controller.Update)
	group.Delete("/:id", controller.Delete)
}
