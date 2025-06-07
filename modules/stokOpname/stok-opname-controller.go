package stokOpname

import (
	"backend/config"
	"backend/models"
	"backend/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service IStokOpnameService
}

func NewStokOpnameController(service IStokOpnameService) *Controller {
	return &Controller{service}
}

func (h *Controller) Index(c *fiber.Ctx) error {
	var stokOpnames []models.StokOpname
	paginated, err := utils.Paginate(
		c,
		config.DB,
		&stokOpnames,
		[]string{"kode_stok_opname", "tanggal", "keterangan", "status"},
		[]string{"tanggal", "lokasi_barang"},
	)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Failed to fetch stock opname data", err.Error())
	}

	return c.JSON(paginated)
}

func (h *Controller) Show(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)

	stokOpname, err := h.service.GetByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Stock opname not found",
		})
	}
	return c.JSON(stokOpname)
}

func (h *Controller) Store(c *fiber.Ctx) error {
	var stokOpname models.StokOpname
	if err := c.BodyParser(&stokOpname); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid data",
		})
	}

	err := h.service.Create(&stokOpname)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to create stock opname",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(stokOpname)
}
