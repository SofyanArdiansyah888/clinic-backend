package produksiBarang

import (
	"backend/config"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(api fiber.Router) {
	repo := NewProduksiBarangRepository(config.DB)
	service := NewProduksiBarangService(repo)
	controller := NewProduksiBarangController(service)

	produksi := api.Group("/produksi-barang")
	produksi.Post("/", controller.Create)
	produksi.Get("/:nomor", controller.GetByNomor)
}
