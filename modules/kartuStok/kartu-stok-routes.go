package kartuStok

import (
	"backend/config"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(api fiber.Router) {
	controller := NewController(NewKartuStokService(NewKartuStokRepository(config.DB)))
	kartuStokRoutes := api.Group("/kartu-stok")
	kartuStokRoutes.Post("/", controller.Show) // Changed to GET since Show is typically for retrieving/viewing data
}
