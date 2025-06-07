package konversiBarang

import (
	"backend/config"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(api fiber.Router) {
	repo := NewKonversiBarangRepository(config.DB)
	service := NewKonversiBarangService(repo)
	controller := NewKonversiBarangController(service)

	konversi := api.Group("/konversi-stok")
	konversi.Post("/", controller.Create)
	konversi.Get("/:nomor", controller.GetByNomor)
}
