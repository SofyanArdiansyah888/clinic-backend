package bank

import (
	"backend/config"
	"backend/models"
	"backend/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service IBankService
}

func NewBankController(service IBankService) *Controller {
	return &Controller{service}
}

func (h *Controller) Index(c *fiber.Ctx) error {
	var banks []models.Bank
	paginated, err := utils.Paginate(
		c,
		config.DB.Preload("Cabang"),
		&banks,
		[]string{"nama_bank", "no_bank", "jenis_bank", "no_rekening", "atas_nama"},
		[]string{"jenis_bank"},
	)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal mengambil data bank", err.Error())
	}
	return c.JSON(paginated)
}

func (h *Controller) Show(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid", err.Error())
	}

	bank, err := h.service.GetByID(uint(id))
	if err != nil {
		return utils.Error(c, fiber.StatusNotFound, "Bank tidak ditemukan", err.Error())
	}
	return c.JSON(bank)
}

func (h *Controller) Store(c *fiber.Ctx) error {
	var bank models.Bank
	if err := c.BodyParser(&bank); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Format data tidak valid", err.Error())
	}

	err := h.service.Create(&bank)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal membuat bank", err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(bank)
}

func (h *Controller) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid", err.Error())
	}

	var bank models.Bank
	if err := c.BodyParser(&bank); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Format data tidak valid", err.Error())
	}
	err = h.service.Update(uint(id), &bank)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal update bank", err.Error())
	}

	return utils.Success(c, "Bank berhasil diupdate", fiber.StatusOK)
}

func (h *Controller) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid", err.Error())
	}

	if err := h.service.Delete(uint(id)); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal hapus bank", err.Error())
	}

	return utils.Success(c, "Bank berhasil dihapus", fiber.StatusOK)
}
