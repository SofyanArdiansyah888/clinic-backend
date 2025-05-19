package routes

import (
	"backend/modules/antrian"
	"backend/modules/bank"
	"backend/modules/cabang"  // Tambahkan import ini
	generateNumber "backend/modules/generateNumber"
	"backend/modules/pasien"
	"backend/modules/perawatan"
	"backend/modules/perusahaan"
	"backend/modules/staff"
	"backend/modules/supplier"
	"backend/modules/treatment"
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
	cabang.RegisterRoutes(api)  // Tambahkan registrasi route cabang
}
