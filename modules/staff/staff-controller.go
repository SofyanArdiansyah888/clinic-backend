package staff

import (
	"backend/config"
	"backend/models"
	"backend/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Controller struct {
	service IStaffService
}

func NewStaffController(service IStaffService) *Controller {
	return &Controller{service}
}

func (h *Controller) Index(c *fiber.Ctx) error {
	var staffs []models.Staff
	paginated, err := utils.Paginate(
		c,
		config.DB,
		&staffs,
		[]string{"nama", "no_staff", "jabatan", "level", "telepon"},
		[]string{"jabatan", "level"},
	)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal mengambil data staff", err.Error())
	}
	return c.JSON(paginated)
}

func (h *Controller) Show(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)
	staff, err := h.service.GetByID(uint(id))
	if err != nil {
		return utils.Error(c, fiber.StatusNotFound, "Staff not found", err.Error())
	}
	return c.JSON(staff)
}

func (h *Controller) Store(c *fiber.Ctx) error {
	var staff models.Staff

	if err := c.BodyParser(&staff); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Data tidak valid", err.Error())
	}

	err := h.service.Create(&staff)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal membuat staff", err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(staff)
}

func (h *Controller) Update(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid")
	}

	var staff models.Staff
	if err := c.BodyParser(&staff); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Data tidak valid")
	}

	if err := h.service.Update(uint(id), &staff); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal update staff", err.Error())
	}
	return utils.Success(c, "Staff berhasil diupdate", 200)
}

func (h *Controller) Delete(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)

	err := h.service.Delete(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal hapus staff",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Staff berhasil dihapus",
	})
}
