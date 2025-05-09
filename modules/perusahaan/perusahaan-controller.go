package perusahaan

import (
	"backend/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Controller struct {
	service IPerusahaanService
}

func NewPerusahaanController(service IPerusahaanService) *Controller {
	return &Controller{service}
}

func (h *Controller) Show(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)

	pasien, err := h.service.GetByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Perusahaan tidak ditemukan",
		})
	}
	return c.JSON(pasien)
}

func (h *Controller) Store(c *fiber.Ctx) error {
	var pasien models.Perusahaan
	if err := c.BodyParser(&pasien); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Data tidak valid",
		})
	}

	err := h.service.Create(&pasien)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal membuat pasien",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(pasien)
}

func (h *Controller) Update(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)

	var pasien models.Perusahaan
	if err := c.BodyParser(&pasien); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Data tidak valid",
		})
	}

	err := h.service.Update(uint(id), &pasien)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal update pasien",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Perusahaan berhasil diupdate",
	})
}
