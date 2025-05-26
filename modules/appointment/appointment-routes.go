package appointment

import (
	"backend/config"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(api fiber.Router) {
	controller := NewAppointmentController(NewAppointmentService(NewAppointmentRepository(config.DB)))
	appointmentRoutes := api.Group("/appointment")
	appointmentRoutes.Get("/", controller.Index)
	appointmentRoutes.Get("/:id", controller.Show)
	appointmentRoutes.Post("/", controller.Store)
	appointmentRoutes.Put("/:id", controller.Update)
	appointmentRoutes.Delete("/:id", controller.Delete)
}
