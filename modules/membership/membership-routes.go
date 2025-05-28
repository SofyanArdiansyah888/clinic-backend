package membership

import (
	"backend/config"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(api fiber.Router) {
	controller := NewMembershipController(NewMembershipService(NewMembershipRepository(config.DB)))
	membershipRoutes := api.Group("/membership")
	membershipRoutes.Get("/", controller.Index)
	membershipRoutes.Get("/:id", controller.Show)
	membershipRoutes.Get("/pasien/:pasien_id", controller.GetByPasien)
	membershipRoutes.Post("/", controller.Store)
	membershipRoutes.Put("/:id", controller.Update)
	membershipRoutes.Put("/:id/points", controller.UpdatePoints)
	membershipRoutes.Delete("/:id", controller.Delete)
}
