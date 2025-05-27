package voucher

import (
	"backend/config"
	"backend/models"
	"backend/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service IVoucherService
}

func NewVoucherController(service IVoucherService) *Controller {
	return &Controller{service}
}

func (h *Controller) Index(c *fiber.Ctx) error {
	var vouchers []models.Voucher
	paginated, err := utils.Paginate(
		c,
		config.DB,
		&vouchers,
		[]string{"kode_voucher", "nama_voucher", "tipe_diskon", "status"},
		[]string{"status"},
	)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal mengambil data voucher", err.Error())
	}

	return c.JSON(paginated)
}

func (h *Controller) Show(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)

	voucher, err := h.service.GetByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Voucher tidak ditemukan",
		})
	}
	return c.JSON(voucher)
}

func (h *Controller) Store(c *fiber.Ctx) error {
	var voucher models.Voucher
	if err := c.BodyParser(&voucher); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Data tidak valid",
		})
	}

	err := h.service.Create(&voucher)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Gagal membuat voucher",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(voucher)
}

func (h *Controller) Update(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)

	var voucher models.Voucher
	if err := c.BodyParser(&voucher); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Data tidak valid",
			"message": err.Error(),
		})
	}

	err := h.service.Update(uint(id), &voucher)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Gagal update voucher",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Voucher berhasil diupdate",
	})
}

func (h *Controller) Delete(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)

	err := h.service.Delete(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Gagal hapus voucher",
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Voucher berhasil dihapus",
	})
}
