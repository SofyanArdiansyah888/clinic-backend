package membership

import (
	"backend/config"
	"backend/models"
	"backend/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service IMembershipService
}

func NewMembershipController(service IMembershipService) *Controller {
	return &Controller{service: service}
}

func (h *Controller) Index(c *fiber.Ctx) error {
	var memberships []models.Membership
	paginated, err := utils.Paginate(
		c,
		config.DB.Preload("Pasien"),
		&memberships,
		[]string{"no_membership", "tipe_membership", "status"},
		[]string{"created_at", "status"},
	)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal mengambil data membership", err.Error())
	}
	return c.JSON(paginated)
}

func (h *Controller) Show(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid", err.Error())
	}

	membership, err := h.service.GetByID(uint(id))
	if err != nil {
		return utils.Error(c, fiber.StatusNotFound, "Membership tidak ditemukan", err.Error())
	}
	return c.JSON(membership)
}

func (h *Controller) GetByPasien(c *fiber.Ctx) error {
	pasienID, err := strconv.Atoi(c.Params("pasien_id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID Pasien tidak valid", err.Error())
	}

	memberships, err := h.service.GetByPasienID(uint(pasienID))
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal mengambil data membership", err.Error())
	}
	return c.JSON(memberships)
}

func (h *Controller) Store(c *fiber.Ctx) error {
	var membership models.Membership
	if err := c.BodyParser(&membership); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Format data tidak valid", err.Error())
	}

	err := h.service.Create(&membership)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal membuat membership", err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(membership)
}

func (h *Controller) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid", err.Error())
	}

	var membership models.Membership
	if err := c.BodyParser(&membership); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Format data tidak valid", err.Error())
	}

	err = h.service.Update(uint(id), &membership)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal update membership", err.Error())
	}

	return utils.Success(c, "Membership berhasil diupdate", fiber.StatusOK)
}

func (h *Controller) UpdatePoints(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid", err.Error())
	}

	type PointsUpdate struct {
		Points int `json:"points"`
	}

	var pointsUpdate PointsUpdate
	if err := c.BodyParser(&pointsUpdate); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Format data tidak valid", err.Error())
	}

	err = h.service.UpdatePoints(uint(id), pointsUpdate.Points)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal update poin membership", err.Error())
	}

	return utils.Success(c, "Poin membership berhasil diupdate", fiber.StatusOK)
}

func (h *Controller) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid", err.Error())
	}

	if err := h.service.Delete(uint(id)); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal hapus membership", err.Error())
	}

	return utils.Success(c, "Membership berhasil dihapus", fiber.StatusOK)
}
