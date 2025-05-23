package user

import (
	"backend/config"
	"backend/models"
	"backend/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	GetAll() ([]models.User, error)
	GetByID(id uint) (*models.User, error)
	Create(data *models.User) error
	Update(id uint, data *models.User) error
	Delete(id uint) error
}

type userService struct {
	repo IUserRepository
}

func NewUserService(repo IUserRepository) *userService {
	return &userService{repo: repo}
}

func (s *userService) GetAll() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetByID(id uint) (*models.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("user tidak ditemukan")
	}
	return user, nil
}

func (s *userService) Create(data *models.User) error {
	// Validasi
	if data.Nama == "" {
		return errors.New("nama user harus diisi")
	}
	if data.Password == "" {
		return errors.New("password harus diisi")
	}

	// Generate nomor user
	data.NoUser = utils.GenerateID(config.DB, "USR", true)

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	data.Password = string(hashedPassword)

	// Cek apakah username telah ada
	existingUser, err := s.repo.FindByUsername(data.Username)
	if err == nil && existingUser != nil {
		return errors.New("Username telah digunakan")
	}

	return s.repo.Create(data)
}

func (s *userService) Update(id uint, data *models.User) error {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("user tidak ditemukan")
	}

	// Validasi
	if data.Nama == "" {
		return errors.New("nama user harus diisi")
	}

	// Jika password diisi, hash password baru
	if data.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		data.Password = string(hashedPassword)
	} else {
		data.Password = user.Password
	}

	// Preserve fields
	data.NoUser = user.NoUser
	data.ID = user.ID

	return s.repo.Update(data)
}

func (s *userService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("user tidak ditemukan")
	}
	return s.repo.Delete(id)
}
