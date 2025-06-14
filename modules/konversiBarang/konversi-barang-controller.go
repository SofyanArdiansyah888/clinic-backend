package konversiBarang

import (
	"backend/config"
	"backend/models"
	"backend/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type KonversiBarangController struct {
	service *KonversiBarangService
}

func NewKonversiBarangController(service *KonversiBarangService) *KonversiBarangController {
	return &KonversiBarangController{service: service}
}

func (h *KonversiBarangController) Index(c *fiber.Ctx) error {
	var konversis []models.KonversiStok

	paginated, err := utils.Paginate(
		c,
		config.DB,
		&konversis,
		[]string{"no_konversi"},
		[]string{},
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch conversion data",
		})
	}

	return c.JSON(paginated)
}

var validate = validator.New()

func (c *KonversiBarangController) Create(ctx *fiber.Ctx) error {
	var req CreateKonversiRequest
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
			"message": "Failed to create konversi barang",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":   "Konversi barang created successfully",
		"transaksi": transaksi,
	})
}

func (c *KonversiBarangController) GetByNomor(ctx *fiber.Ctx) error {
	nomorTransaksi := ctx.Params("nomor")

	transaksi, details, err := c.service.GetByNomor(nomorTransaksi)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Konversi barang not found",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"transaksi": transaksi,
		"details":   details,
	})
}
