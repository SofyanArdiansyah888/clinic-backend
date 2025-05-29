package pembelianBarang

import (
	"github.com/gofiber/fiber/v2"
	"backend/config"
)

func RegisterRoutes(api fiber.Router) {
	repo := NewPembelianBarangRepository(config.DB)
	service := NewPembelianBarangService(repo)
	controller := NewPembelianBarangController(service)

	pembelian := api.Group("/pembelian")
	pembelian.Post("/", controller.Create)
	pembelian.Get("/:nomor", controller.GetByNomor)
}