package promo

import (
	"backend/models"
	"backend/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PromoController struct {
	service *PromoService
}

func NewPromoController(service *PromoService) *PromoController {
	return &PromoController{service}
}

func (h *PromoController) Index(c *fiber.Ctx) error {
	promos, err := h.service.GetAll()
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal mengambil data promo", err.Error())
	}
	return c.JSON(promos)
}

func (h *PromoController) Show(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid", err.Error())
	}

	promo, err := h.service.GetByID(uint(id))
	if err != nil {
		return utils.Error(c, fiber.StatusNotFound, "Promo tidak ditemukan", err.Error())
	}
	return c.JSON(promo)
}

func (h *PromoController) Store(c *fiber.Ctx) error {
	var promo models.Promo
	if err := c.BodyParser(&promo); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Data tidak valid", err.Error())
	}

	if err := h.service.Create(&promo); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal membuat promo", err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(promo)
}

func (h *PromoController) Update(c *fiber.Ctx) error {
	_, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid", err.Error())
	}

	var promo models.Promo
	if err := c.BodyParser(&promo); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Data tidak valid", err.Error())
	}

	if err := h.service.Update(&promo); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal update promo", err.Error())
	}

	return utils.Success(c, "Promo berhasil diupdate", fiber.StatusOK)
}

func (h *PromoController) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid", err.Error())
	}

	if err := h.service.Delete(uint(id)); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal hapus promo", err.Error())
	}

	return utils.Success(c, "Promo berhasil dihapus", fiber.StatusOK)
}
