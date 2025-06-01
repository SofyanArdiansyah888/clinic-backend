package penjualanBarang

import (
	"backend/config"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(api fiber.Router) {
	repo := NewPenjualanBarangRepository(config.DB)
	service := NewPenjualanBarangService(repo)
	controller := NewPenjualanBarangController(service)

	penjualan := api.Group("/penjualan")
	penjualan.Post("/", controller.Create)
	penjualan.Get("/:nomor", controller.GetByNomor)
}
