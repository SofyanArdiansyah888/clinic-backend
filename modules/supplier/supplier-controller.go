package supplier

import (
	"backend/config"
	"backend/models"
	"backend/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service ISupplierService
}

func NewSupplierController(service ISupplierService) *Controller {
	return &Controller{service}
}

func (h *Controller) Index(c *fiber.Ctx) error {
	var suppliers []models.Supplier
	paginated, err := utils.Paginate(
		c,
		config.DB,
		&suppliers,
		[]string{"nama", "no_supplier", "telepon", "alamat"},
		[]string{"nama"},
	)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal mengambil data supplier", err.Error())
	}
	return c.JSON(paginated)
}

func (h *Controller) Show(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid", err.Error())
	}

	supplier, err := h.service.GetByID(uint(id))
	if err != nil {
		return utils.Error(c, fiber.StatusNotFound, "Supplier tidak ditemukan", err.Error())
	}
	return c.JSON(supplier)
}

func (h *Controller) Store(c *fiber.Ctx) error {
	var supplier models.Supplier
	if err := c.BodyParser(&supplier); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Data tidak valid", err.Error())
	}

	err := h.service.Create(&supplier)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Gagal membuat supplier",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(supplier)
}

func (h *Controller) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid", err.Error())
	}

	var supplier models.Supplier
	if err := c.BodyParser(&supplier); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Data tidak valid", err.Error())
	}

	if err := h.service.Update(uint(id), &supplier); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal update supplier", err.Error())
	}

	return utils.Success(c, "Supplier berhasil diupdate", fiber.StatusOK)
}

func (h *Controller) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid", err.Error())
	}

	if err := h.service.Delete(uint(id)); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal hapus supplier", err.Error())
	}

	return utils.Success(c, "Supplier berhasil dihapus", fiber.StatusOK)
}
