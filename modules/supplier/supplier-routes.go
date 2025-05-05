package supplier

import (
	"backend/config"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(api fiber.Router) {
	controller := NewSupplierController(NewSupplierService(NewSupplierRepository(config.DB)))
	supplierRoutes := api.Group("/supplier")
	supplierRoutes.Get("/", controller.Index)
	supplierRoutes.Get("/:id", controller.Show)
	supplierRoutes.Post("/", controller.Store)
	supplierRoutes.Put("/:id", controller.Update)
	supplierRoutes.Delete("/:id", controller.Delete)
}
