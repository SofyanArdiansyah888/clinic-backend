package antrian

import (
	"backend/config"
	"backend/models"
	"backend/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service IAntrianService
}

func NewAntrianController(service IAntrianService) *Controller {
	return &Controller{service}
}

func (h *Controller) Index(c *fiber.Ctx) error {
	var antrians []models.Antrian
	query := config.DB.Preload("Pasien").Preload("Staff")
	paginated, err := utils.Paginate(
		c,
		query,
		&antrians,
		[]string{"tanggal"},
		[]string{"status"},
	)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal mengambil data antrian", err.Error())
	}

	return c.JSON(paginated)
}

func (h *Controller) Show(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)

	antrian, err := h.service.GetByID(uint(id))
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Antrian tidak ditemukan", err.Error())
	}
	return utils.Success(c, "Antrian ditemukan", 200, antrian)
}

func (h *Controller) CreateUserAntrian(c *fiber.Ctx) error {
	var pasien models.Pasien
	if err := c.BodyParser(&pasien); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Data pasien tidak valid", err.Error())
	}

	var antrian models.Antrian
	antrian.NoAntrian = utils.GenerateID(config.DB, "ANT", true)
	antrian.IDPasien = pasien.ID
	if err := c.BodyParser(&antrian); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Data antrian tidak valid", err.Error())
	}

	err := h.service.createPasienAntrian(&pasien, &antrian)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal membuat antrian", err.Error())
	}

	return utils.Success(c, "Pasien dan Antrian berhasil dibuat", 200)
}

func (h *Controller) Store(c *fiber.Ctx) error {
	var antrian models.Antrian
	if err := c.BodyParser(&antrian); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Data tidak valid")
	}
	antrian.NoAntrian = utils.GenerateID(config.DB, "ANT", true)
	err := h.service.Create(&antrian)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal membuat antrian", err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(antrian)
}

func (h *Controller) Update(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)

	var antrian models.Antrian
	if err := c.BodyParser(&antrian); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Data tidak valid",
		})
	}

	err := h.service.Update(uint(id), &antrian)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal update antrian",
		})
	}

	return utils.Success(c, "Antrian berhasil diupdate", 200)
}

func (h *Controller) Delete(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)

	err := h.service.Delete(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal hapus antrian",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Antrian berhasil dihapus",
	})
}
