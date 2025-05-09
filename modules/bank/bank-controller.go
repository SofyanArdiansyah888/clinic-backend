package bank

import (
	"backend/config"
	"backend/models"
	"backend/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Controller struct {
	service IBankService
}

func NewBankController(service IBankService) *Controller {
	return &Controller{service}
}

func (h *Controller) Index(c *fiber.Ctx) error {
	var banks []models.Bank

	//db := h.service.(*bankService).repo.GetDB()

	paginated, err := utils.Paginate(
		c,
		config.DB,
		&banks,
		[]string{"nama", "no_bank", "telepon", "alamat"},
		[]string{},
	)
	if err != nil {
		return utils.Error(c, 500, "Gagal mengambil data bank", err.Error())
	}

	return c.JSON(paginated)
}

func (h *Controller) Show(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)

	bank, err := h.service.GetByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Bank tidak ditemukan",
		})
	}
	return c.JSON(bank)
}

func (h *Controller) Store(c *fiber.Ctx) error {
	var bank models.Bank
	if err := c.BodyParser(&bank); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Data tidak valid",
		})
	}

	err := h.service.Create(&bank)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal membuat bank",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(bank)
}

func (h *Controller) Update(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)

	var bank models.Bank
	if err := c.BodyParser(&bank); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Data tidak valid",
		})
	}

	err := h.service.Update(uint(id), &bank)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal update bank",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Bank berhasil diupdate",
	})
}

func (h *Controller) Delete(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)

	err := h.service.Delete(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal hapus bank",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Bank berhasil dihapus",
	})
}
