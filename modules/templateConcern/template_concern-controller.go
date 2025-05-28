package templateConcern

import (
	"backend/config"
	"backend/models"
	"backend/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TemplateConcernController struct {
	service *TemplateConcernService
}

func NewTemplateConcernController(service *TemplateConcernService) *TemplateConcernController {
	return &TemplateConcernController{service}
}

func (h *TemplateConcernController) Index(c *fiber.Ctx) error {
	var concerns []models.TemplateConcern
	paginated, err := utils.Paginate(
		c,
		config.DB,
		&concerns,
		[]string{"nama", "deskripsi"},
		[]string{"nama"},
	)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal mengambil data template concern", err.Error())
	}
	return c.JSON(paginated)
}

func (h *TemplateConcernController) Show(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid", err.Error())
	}

	concern, err := h.service.GetByID(uint(id))
	if err != nil {
		return utils.Error(c, fiber.StatusNotFound, "Template concern tidak ditemukan", err.Error())
	}
	return c.JSON(concern)
}

func (h *TemplateConcernController) Store(c *fiber.Ctx) error {
	var concern models.TemplateConcern
	if err := c.BodyParser(&concern); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Data tidak valid", err.Error())
	}

	if err := h.service.Create(&concern); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal membuat template concern", err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(concern)
}

func (h *TemplateConcernController) Update(c *fiber.Ctx) error {
	_, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid", err.Error())
	}

	var concern models.TemplateConcern
	if err := c.BodyParser(&concern); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Data tidak valid", err.Error())
	}

	if err := h.service.Update(&concern); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal update template concern", err.Error())
	}

	return utils.Success(c, "Template concern berhasil diupdate", fiber.StatusOK)
}

func (h *TemplateConcernController) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid", err.Error())
	}

	if err := h.service.Delete(uint(id)); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal hapus template concern", err.Error())
	}

	return utils.Success(c, "Template concern berhasil dihapus", fiber.StatusOK)
}
