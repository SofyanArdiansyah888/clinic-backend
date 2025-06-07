package kartuStok

import (
	"backend/utils"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service IKartuStokService
}

func NewController(service IKartuStokService) *Controller {
	return &Controller{service: service}
}

func (h *Controller) Show(c *fiber.Ctx) error {
	var request struct {
		KodeBarang string `json:"kode_barang"`
		Dari       string `json:"dari"`
		Sampai     string `json:"sampai"`
	}
	if err := c.BodyParser(&request); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}
	kodeBarang := request.KodeBarang
	if kodeBarang == "" {
		return utils.Error(c, fiber.StatusBadRequest, "Kode barang is required", kodeBarang)
	}

	movementStok, err := h.service.GetKartuStok(kodeBarang, request.Dari, request.Sampai)
	if err != nil {
		return utils.Error(c, fiber.StatusNotFound, "Barang tidak ditemukan", err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(movementStok)
}
