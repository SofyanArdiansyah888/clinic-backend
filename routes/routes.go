package routes

import (
	"backend/modules/antrian"
	generateNumber "backend/modules/generateNumber"
	"backend/modules/pasien"
	"backend/modules/perawatan"
	"backend/modules/staff"
	"backend/modules/supplier"
	"backend/modules/treatment"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	pasien.RegisterRoutes(api)
	antrian.RegisterRoutes(api)
	staff.RegisterRoutes(api)
	generateNumber.RegisterRoutes(api)
	treatment.RegisterRoutes(api)
	supplier.RegisterRoutes(api)
	perawatan.RegisterRoutes(api)
}
