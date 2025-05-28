package appointment

import (
	"backend/config"
	"backend/models"
	"backend/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service IAppointmentService
}

func NewAppointmentController(service IAppointmentService) *Controller {
	return &Controller{service: service}
}

func (h *Controller) Index(c *fiber.Ctx) error {
	var appointments []models.Appointment
	paginated, err := utils.Paginate(
		c,
		config.DB.Preload("Pasien").Preload("Staff").Preload("Cabang"),
		&appointments,
		[]string{"tanggal", "jam_mulai", "jam_selesai", "status"},
		[]string{"tanggal", "status"},
	)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal mengambil data appointment", err.Error())
	}
	return c.JSON(paginated)
}

func (h *Controller) Show(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid", err.Error())
	}

	appointment, err := h.service.GetByID(uint(id))
	if err != nil {
		return utils.Error(c, fiber.StatusNotFound, "Appointment tidak ditemukan", err.Error())
	}
	return c.JSON(appointment)
}

func (h *Controller) Store(c *fiber.Ctx) error {
	var appointment models.Appointment
	if err := c.BodyParser(&appointment); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Format data tidak valid", err.Error())
	}

	err := h.service.Create(&appointment)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal membuat appointment", err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(appointment)
}

func (h *Controller) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid", err.Error())
	}

	var appointment models.Appointment
	if err := c.BodyParser(&appointment); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Format data tidak valid", err.Error())
	}

	err = h.service.Update(uint(id), &appointment)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal update appointment", err.Error())
	}

	return utils.Success(c, "Appointment berhasil diupdate", fiber.StatusOK)
}

func (h *Controller) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid", err.Error())
	}

	if err := h.service.Delete(uint(id)); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal hapus appointment", err.Error())
	}

	return utils.Success(c, "Appointment berhasil dihapus", fiber.StatusOK)
}
