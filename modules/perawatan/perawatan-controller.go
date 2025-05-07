package perawatan

import (
	"backend/config"
	"backend/models"
	"backend/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Controller struct {
	service IPerawatanService
}

func NewPerawatanController(service IPerawatanService) *Controller {
	return &Controller{service}
}

func (h *Controller) Index(c *fiber.Ctx) error {
	var pasiens []models.Perawatan

	//db := h.service.(*pasienService).repo.GetDB()

	paginated, err := utils.Paginate(
		c,
		config.DB,
		&pasiens,
		[]string{"nama_pasien", "no_rm", "no_member"},
		[]string{},
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal mengambil data pasien",
		})
	}

	return c.JSON(paginated)
}

func (h *Controller) Show(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)

	pasien, err := h.service.GetByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Perawatan tidak ditemukan",
		})
	}
	return c.JSON(pasien)
}

func (h *Controller) Store(c *fiber.Ctx) error {
	var pasien models.Perawatan
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

	var pasien models.Perawatan
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
		"message": "Perawatan berhasil diupdate",
	})
}

func (h *Controller) Delete(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)

	err := h.service.Delete(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal hapus pasien",
		})
	}

	return c.JSON(fiber.Map{
		"message": " Perawatan berhasil dihapus",
	})
}
