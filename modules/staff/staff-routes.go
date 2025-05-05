package staff

import (
	"backend/config"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(api fiber.Router) {
	repository := NewStaffRepository(config.DB)
	service := NewStaffService(repository)
	controller := NewStaffController(service)
	group := api.Group("/staff")
	group.Get("/", controller.Index)
	group.Get("/:id", controller.Show)
	group.Post("/", controller.Store)
	group.Put("/:id", controller.Update)
	group.Delete("/:id", controller.Delete)

}
