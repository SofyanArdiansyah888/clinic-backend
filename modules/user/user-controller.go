package user

import (
	"backend/config"
	"backend/models"
	"backend/utils"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service IUserService
}

func NewUserController(service IUserService) *Controller {
	return &Controller{service}
}

func (h *Controller) Index(c *fiber.Ctx) error {
	var users []models.User
	query := config.DB.Preload("Cabang")
	paginated, err := utils.Paginate(
		c,
		query,
		&users,
		[]string{"nama", "username", "role"},
		[]string{"nama"},
	)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal mengambil data user", err.Error())
	}
	return c.JSON(paginated)
}

func (h *Controller) Show(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid", err.Error())
	}

	user, err := h.service.GetByID(uint(id))
	if err != nil {
		return utils.Error(c, fiber.StatusNotFound, "User tidak ditemukan", err.Error())
	}
	return c.JSON(user)
}

func (h *Controller) Store(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Data tidak valid", err.Error())
	}

	if err := h.service.Create(&user); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal membuat user", err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *Controller) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid", err.Error())
	}

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Data tidak valid", err.Error())
	}

	if err := h.service.Update(uint(id), &user); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal update user", err.Error())
	}

	return utils.Success(c, "User berhasil diupdate", fiber.StatusOK)
}

func (h *Controller) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "ID tidak valid", err.Error())
	}

	if err := h.service.Delete(uint(id)); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "Gagal hapus user", err.Error())
	}

	return utils.Success(c, "User berhasil dihapus", fiber.StatusOK)
}

func SeedUsers(c *fiber.Ctx) error {
	for i := 1; i <= 30; i++ {
		user := models.User{
			Nama:     fmt.Sprintf("User %d", i),
			Username: fmt.Sprintf("user%d", i),
			Role:     "user",
			IDCabang: 1,
			NoUser:   utils.GenerateID(config.DB, "USR", true),
			Password: "password", // use a hashed password in a real application
		}

		if err := config.DB.Create(&user).Error; err != nil {
			return utils.Error(c, fiber.StatusInternalServerError, "Gagal seeding data user", err.Error())
		}
	}
	return utils.Success(c, "Berhasil seeding 30 user", fiber.StatusOK)
}
