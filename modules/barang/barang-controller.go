package barang

import (
	transaksibarang "backend/modules/transaksiBarang"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	BarangRepo          IBarangRepository
	TransaksiBarangRepo transaksibarang.ITransaksiBarangRepository
}

func NewController(
	barangRepo IBarangRepository,
	transaksibarangRepo transaksibarang.ITransaksiBarangRepository,
) *Controller {
	return &Controller{
		BarangRepo:          barangRepo,
		TransaksiBarangRepo: transaksibarangRepo,
	}
}

func (h *Controller) Index(c *fiber.Ctx) error {
	barang, err := h.BarangRepo.GetBarang()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data barang"})
	}
	return c.JSON(barang)
}

func (h *Controller) Store(c *fiber.Ctx) error {
	_, err := h.TransaksiBarangRepo.CreateTransaksi()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data barang"})
	}
	return c.JSON(fiber.Map{"message": "Hello world"})
}
