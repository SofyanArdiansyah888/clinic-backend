package penjualanBarang

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type PenjualanBarangController struct {
	service *PenjualanBarangService
}

func NewPenjualanBarangController(service *PenjualanBarangService) *PenjualanBarangController {
	return &PenjualanBarangController{service: service}
}

var validate = validator.New()

func (c *PenjualanBarangController) Create(ctx *fiber.Ctx) error {
	var req CreatePenjualanRequest
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
			"message": "Failed to create penjualan",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":   "Penjualan created successfully",
		"transaksi": transaksi,
	})
}

func (c *PenjualanBarangController) GetByNomor(ctx *fiber.Ctx) error {
	nomorTransaksi := ctx.Params("nomor")

	transaksi, details, err := c.service.GetByNomor(nomorTransaksi)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Penjualan not found",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"transaksi": transaksi,
		"details":   details,
	})
}
