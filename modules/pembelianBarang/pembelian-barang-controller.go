package pembelianBarang

import (
	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
)

type PembelianBarangController struct {
	service *PembelianBarangService
}

func NewPembelianBarangController(service *PembelianBarangService) *PembelianBarangController {
	return &PembelianBarangController{service: service}
}

var validate = validator.New()

func (c *PembelianBarangController) Create(ctx *fiber.Ctx) error {
	var req CreatePembelianRequest
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

	transaksi, details, err := c.service.Create(req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create pembelian",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":  "Pembelian created successfully",
		"transaksi": transaksi,
		"details":   details,
	})
}

func (c *PembelianBarangController) GetByNomor(ctx *fiber.Ctx) error {
	nomorTransaksi := ctx.Params("nomor")

	transaksi, details, err := c.service.GetByNomor(nomorTransaksi)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Pembelian not found",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"transaksi": transaksi,
		"details":   details,
	})
}