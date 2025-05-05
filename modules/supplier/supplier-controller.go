package supplier

import (
	"backend/config"
	"backend/models"
	"backend/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Controller struct {
	service ISupplierService
}

func NewSupplierController(service ISupplierService) *Controller {
	return &Controller{service}
}

func (h *Controller) Index(c *fiber.Ctx) error {
	var suppliers []models.Supplier

	//db := h.service.(*supplierService).repo.GetDB()

	paginated, err := utils.Paginate(
		c,
		config.DB,
		&suppliers,
		[]string{"nama", "no_supplier", "telepon", "alamat"},
		[]string{},
	)
	if err != nil {
		return utils.Error(c, 500, "Gagal mengambil data supplier", err.Error())
	}

	return c.JSON(paginated)
}

func (h *Controller) Show(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)

	supplier, err := h.service.GetByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Supplier tidak ditemukan",
		})
	}
	return c.JSON(supplier)
}

func (h *Controller) Store(c *fiber.Ctx) error {
	var supplier models.Supplier
	if err := c.BodyParser(&supplier); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Data tidak valid",
		})
	}

	err := h.service.Create(&supplier)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal membuat supplier",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(supplier)
}

func (h *Controller) Update(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)

	var supplier models.Supplier
	if err := c.BodyParser(&supplier); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Data tidak valid",
		})
	}

	err := h.service.Update(uint(id), &supplier)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal update supplier",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Supplier berhasil diupdate",
	})
}

func (h *Controller) Delete(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)

	err := h.service.Delete(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal hapus supplier",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Supplier berhasil dihapus",
	})
}
