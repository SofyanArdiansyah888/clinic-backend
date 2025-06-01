package stokOpname

import (
	"backend/config"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(api fiber.Router) {
	repository := NewStokOpnameRepository(config.DB)
	service := NewStokOpnameService(repository)
	controller := NewStokOpnameController(service)

	group := api.Group("/stok-opname")
	group.Get("/", controller.Index)
	group.Get("/:id", controller.Show)
	group.Post("/", controller.Store)
}
