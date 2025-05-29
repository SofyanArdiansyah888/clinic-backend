package pembelianBarang

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
	repo := NewPembelianBarangRepository(db)
	service := NewPembelianBarangService(repo)
	controller := NewPembelianBarangController(service)

	pembelian := api.Group("/pembelian")
	pembelian.Post("/", controller.Create)
	pembelian.Get("/:nomor", controller.GetByNomor)
}