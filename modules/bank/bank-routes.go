package bank

import (
	"backend/config"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(api fiber.Router) {
	controller := NewBankController(NewBankService(NewBankRepository(config.DB)))
	bankRoutes := api.Group("/bank")
	bankRoutes.Get("/", controller.Index)
	bankRoutes.Get("/:id", controller.Show)
	bankRoutes.Post("/", controller.Store)
	bankRoutes.Put("/:id", controller.Update)
	bankRoutes.Delete("/:id", controller.Delete)
}
