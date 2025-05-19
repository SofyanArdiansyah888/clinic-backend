package bank

import (
	"backend/config"
	"backend/models"
	"backend/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service IBankService
}

func NewBankController(service IBankService) *Controller {
	return &Controller{service}
}

func (h *Controller) Index(c *fiber.Ctx) error {
	var banks []models.Bank
	paginated, err := utils.Paginate(
		c,
		config.DB.Preload("Cabang"),
		&banks,
		[]string{"nama_bank", "no_bank", "jenis_bank"},
		[]string{"jenis_bank"},
	)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal mengambil data bank", err.Error())
	}
	return c.JSON(paginated)
}

func (h *Controller) Show(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid", err.Error())
	}

	bank, err := h.service.GetByID(uint(id))
	if err != nil {
		return utils.Error(c, fiber.StatusNotFound, "Bank tidak ditemukan", err.Error())
	}
	return c.JSON(bank)
}

func (h *Controller) Store(c *fiber.Ctx) error {
	var bank models.Bank

	// Parse body JSON
	if err := c.BodyParser(&bank); err != nil {
		// Coba parse id_cabang dari form atau JSON string
		idCabangStr := c.FormValue("id_cabang")
		if idCabangStr == "" {
			// Coba ambil dari body JSON
			body := make(map[string]interface{})
			if err := c.BodyParser(&body); err == nil {
				if idCabang, ok := body["id_cabang"].(string); ok {
					idCabangStr = idCabang
				}
			}
		}

		// Konversi id_cabang ke uint
		if idCabangStr != "" {
			idCabang, err := strconv.ParseUint(idCabangStr, 10, 32)
			if err != nil {
				return utils.Error(c, fiber.StatusBadRequest, "ID Cabang harus berupa angka valid", err.Error())
			}
			bank.IDCabang = uint(idCabang)
		}
	}

	if bank.IDCabang == 0 {
		return utils.Error(c, fiber.StatusBadRequest, "ID Cabang harus diisi", "")
	}

	err := h.service.Create(&bank)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal membuat bank", err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(bank)
}

func (h *Controller) Update(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid", err.Error())
	}

	var bank models.Bank

	// Parse body JSON
	if err := c.BodyParser(&bank); err != nil {
		// Coba parse id_cabang dari form atau JSON string
		idCabangStr := c.FormValue("id_cabang")
		if idCabangStr == "" {
			// Coba ambil dari body JSON
			body := make(map[string]interface{})
			if err := c.BodyParser(&body); err == nil {
				if idCabang, ok := body["id_cabang"].(string); ok {
					idCabangStr = idCabang
				}
			}
		}

		// Konversi id_cabang ke uint
		if idCabangStr != "" {
			idCabang, err := strconv.ParseUint(idCabangStr, 10, 32)
			if err != nil {
				return utils.Error(c, fiber.StatusBadRequest, "ID Cabang harus berupa angka valid", err.Error())
			}
			bank.IDCabang = uint(idCabang)
		}
	}

	if bank.IDCabang == 0 {
		return utils.Error(c, fiber.StatusBadRequest, "ID Cabang harus diisi", "")
	}

	err = h.service.Update(uint(id), &bank)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal update bank", err.Error())
	}

	return utils.Success(c, "Bank berhasil diupdate", fiber.StatusOK)
}

func (h *Controller) Delete(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, _ := strconv.Atoi(idParam)

	err := h.service.Delete(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal hapus bank",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Bank berhasil dihapus",
	})
}
