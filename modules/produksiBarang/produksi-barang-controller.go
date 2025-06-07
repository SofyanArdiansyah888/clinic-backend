package produksiBarang

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ProduksiBarangController struct {
	service *ProduksiBarangService
}

func NewProduksiBarangController(service *ProduksiBarangService) *ProduksiBarangController {
	return &ProduksiBarangController{service: service}
}

var validate = validator.New()

func (c *ProduksiBarangController) Create(ctx *fiber.Ctx) error {
	var req CreateProduksiRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	// Validasi request
	if err := validate.Struct(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}

	transaksi, err := c.service.Create(req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create produksi barang",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":   "Produksi barang created successfully",
		"transaksi": transaksi,
	})
}

func (c *ProduksiBarangController) GetByNomor(ctx *fiber.Ctx) error {
	nomorTransaksi := ctx.Params("nomor")

	transaksi, details, err := c.service.GetByNomor(nomorTransaksi)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Produksi barang not found",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"transaksi": transaksi,
		"details":   details,
	})
}
