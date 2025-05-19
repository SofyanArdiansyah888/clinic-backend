package cabang

import (
	"backend/config"
	"backend/models"
	"backend/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Controller struct {
	service ICabangService
}

func NewCabangController(service ICabangService) *Controller {
	return &Controller{service}
}

func (h *Controller) Index(c *fiber.Ctx) error {
	var cabangs []models.Cabang
	paginated, err := utils.Paginate(
		c,
		config.DB,
		&cabangs,
		[]string{"no_cabang", "nama_klinik", "alamat_lengkap", "email_klinik"},
		[]string{"nama_klinik"},
	)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal mengambil data cabang", err.Error())
	}

	return c.JSON(paginated)
}

func (h *Controller) Show(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)

	cabang, err := h.service.GetByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Cabang tidak ditemukan",
		})
	}
	return c.JSON(cabang)
}

func (h *Controller) Store(c *fiber.Ctx) error {
	var cabang models.Cabang
	if err := c.BodyParser(&cabang); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Data tidak valid",
		})
	}

	err := h.service.Create(&cabang)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Gagal membuat cabang",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(cabang)
}

func (h *Controller) Update(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)

	var cabang models.Cabang
	if err := c.BodyParser(&cabang); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Data tidak valid",
			"message": err.Error(),
		})
	}

	err := h.service.Update(uint(id), &cabang)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Gagal update cabang",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Cabang berhasil diupdate",
	})
}

func (h *Controller) Delete(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)

	err := h.service.Delete(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Gagal hapus cabang",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Cabang berhasil dihapus",
	})
}