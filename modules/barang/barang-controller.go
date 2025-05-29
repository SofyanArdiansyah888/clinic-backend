package barang

import (
	"backend/config"
	"backend/models"
	"backend/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service IBarangService
}

func NewController(service IBarangService) *Controller {
	return &Controller{service: service}
}

func (h *Controller) Index(c *fiber.Ctx) error {
	var barangs []models.Barang
	paginated, err := utils.Paginate(
		c,
		config.DB,
		&barangs,
		[]string{"nama_barang", "kode_barang"},
		[]string{"satuan","jenis_barang"},
	)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal mengambil data barang", err.Error())
	}
	return c.JSON(paginated)
}

func (h *Controller) Show(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid", err.Error())
	}

	barang, err := h.service.GetByID(uint(id))
	if err != nil {
		return utils.Error(c, fiber.StatusNotFound, "Barang tidak ditemukan", err.Error())
	}
	return c.JSON(barang)
}

func (h *Controller) Store(c *fiber.Ctx) error {
	var barang models.Barang
	if err := c.BodyParser(&barang); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Format data tidak valid", err.Error())
	}

	err := h.service.Create(&barang)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal membuat barang", err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(barang)
}

func (h *Controller) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid", err.Error())
	}

	var barang models.Barang
	if err := c.BodyParser(&barang); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Format data tidak valid", err.Error())
	}

	err = h.service.Update(uint(id), &barang)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal update barang", err.Error())
	}

	return utils.Success(c, "Barang berhasil diupdate", fiber.StatusOK)
}

func (h *Controller) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid", err.Error())
	}

	if err := h.service.Delete(uint(id)); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal hapus barang", err.Error())
	}

	return utils.Success(c, "Barang berhasil dihapus", fiber.StatusOK)
}
