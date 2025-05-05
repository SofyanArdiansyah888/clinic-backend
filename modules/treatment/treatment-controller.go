package treatment

import (
	"backend/config"
	"backend/models"
	"backend/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Controller struct {
	service ITreatmentService
}

func NewTreatmentController(service ITreatmentService) *Controller {
	return &Controller{service}
}

func (h *Controller) Index(c *fiber.Ctx) error {
	var treatments []models.Treatment
	paginated, err := utils.Paginate(
		c,
		config.DB,
		&treatments,
		[]string{"nama", "alamat", "no_treatment", "telepon", "jabatan", "level"},
		[]string{"jabatan", "level"},
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal mengambil data treatment",
		})
	}

	return c.JSON(paginated)
}

func (h *Controller) Show(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)

	treatment, err := h.service.GetByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Treatment tidak ditemukan",
		})
	}
	return c.JSON(treatment)
}

func (h *Controller) Store(c *fiber.Ctx) error {
	var treatment models.Treatment
	if err := c.BodyParser(&treatment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Data tidak valid",
		})
	}

	err := h.service.Create(&treatment)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Gagal membuat treatment",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(treatment)
}

func (h *Controller) Update(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)

	var treatment models.Treatment
	if err := c.BodyParser(&treatment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Data tidak valid",
			"message": err.Error(),
		})
	}

	err := h.service.Update(uint(id), &treatment)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Gagal update treatment",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Treatment berhasil diupdate",
	})
}

func (h *Controller) Delete(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)

	err := h.service.Delete(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Gagal hapus treatment",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Treatment berhasil dihapus",
	})
}
