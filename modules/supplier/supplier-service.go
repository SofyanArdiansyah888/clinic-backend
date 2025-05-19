package supplier

import (
	"backend/config"
	"backend/models"
	"backend/utils"
	"errors"
	"fmt"
)

type ISupplierService interface {
	GetAll() ([]models.Supplier, error)
	GetByID(id uint) (*models.Supplier, error)
	Create(data *models.Supplier) error
	Update(id uint, data *models.Supplier) error
	Delete(id uint) error
}

type supplierService struct {
	repo ISupplierRepository
}

func NewSupplierService(repo ISupplierRepository) *supplierService {
	return &supplierService{repo: repo}
}

func (s *supplierService) GetAll() ([]models.Supplier, error) {
	return s.repo.FindAll()
}

func (s *supplierService) GetByID(id uint) (*models.Supplier, error) {
	supplier, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("supplier tidak ditemukan")
	}
	return supplier, nil
}

func (s *supplierService) Create(data *models.Supplier) error {
	// Generate unique supplier number
	data.NoSupplier = utils.GenerateID(config.DB, "SUP", true)

	// Validate required fields
	if data.Nama == "" {
		return errors.New("nama supplier harus diisi")
	}
	if data.Alamat == "" {
		return errors.New("alamat supplier harus diisi")
	}
	if data.Telepon == "" {
		return errors.New("telepon supplier harus diisi")
	}
	fmt.Println(data)
	return s.repo.Create(data)
}

func (s *supplierService) Update(id uint, data *models.Supplier) error {
	supplier, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("supplier tidak ditemukan")
	}

	// Validate required fields
	if data.Nama == "" {
		return errors.New("nama supplier harus diisi")
	}
	if data.Alamat == "" {
		return errors.New("alamat supplier harus diisi")
	}
	if data.Telepon == "" {
		return errors.New("telepon supplier harus diisi")
	}

	// Preserve immutable fields
	data.NoSupplier = supplier.NoSupplier
	data.ID = supplier.ID

	return s.repo.Update(data)
}

func (s *supplierService) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("supplier tidak ditemukan")
	}
	return s.repo.Delete(id)
}
