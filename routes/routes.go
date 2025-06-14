package routes

import (
	"backend/modules/antrian"
	"backend/modules/appointment"
	"backend/modules/bank"
	"backend/modules/barang"
	"backend/modules/cabang"
	generateNumber "backend/modules/generateNumber"
	"backend/modules/kartuStok"
	"backend/modules/konversiBarang"
	"backend/modules/membership"
	"backend/modules/pasien"
	"backend/modules/pembelianBarang"
	"backend/modules/penjualanBarang"
	"backend/modules/perawatan"
	"backend/modules/perusahaan"
	"backend/modules/produksiBarang"
	"backend/modules/promo"
	"backend/modules/staff"
	"backend/modules/stokOpname"
	"backend/modules/supplier"
	"backend/modules/templateConcern"
	"backend/modules/treatment"
	"backend/modules/user"
	"backend/modules/voucher"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// SETUP STATIC FILE SERVING
	//app.Static("/", "./tmp")
	//app.Use(func(c *fiber.Ctx) error {
	//	return c.SendFile("./tmp/index.html")
	//})

	// SETUP API ROUTES
	api := app.Group("/api")
	pasien.RegisterRoutes(api)
	antrian.RegisterRoutes(api)
	staff.RegisterRoutes(api)
	generateNumber.RegisterRoutes(api)
	treatment.RegisterRoutes(api)
	supplier.RegisterRoutes(api)
	perawatan.RegisterRoutes(api)
	perusahaan.RegisterRoutes(api)
	bank.RegisterRoutes(api)
	cabang.RegisterRoutes(api)
	user.RegisterRoutes(api)            // Tambahkan registrasi route user
	promo.RegisterRoutes(api)           // Tambahkan registrasi route promo
	templateConcern.RegisterRoutes(api) // Tambahkan registrasi route template concern
	appointment.RegisterRoutes(api)     // Tambahkan registrasi route appointment
	membership.RegisterRoutes(api)      // Tambahkan registrasi route membership
	voucher.RegisterRoutes(api)         // Tambahkan registrasi route voucher
	barang.RegisterRoutes(api)          // Tambahkan registrasi route barang
	pembelianBarang.RegisterRoutes(api) // Tambahkan registrasi route pembelian
	penjualanBarang.RegisterRoutes(api)
	stokOpname.RegisterRoutes(api)
	kartuStok.RegisterRoutes(api)
	konversiBarang.RegisterRoutes(api)
	produksiBarang.RegisterRoutes(api)
}
