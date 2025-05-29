package barang

import (
	"backend/config"
	"backend/models"
	"backend/utils"
)

type IBarangService interface {
	GetAll() ([]models.Barang, error)
	GetByID(id uint) (*models.Barang, error)
	Create(data *models.Barang) error
	Update(id uint, data *models.Barang) error
	Delete(id uint) error
}

type barangService struct {
	repo IBarangRepository
}

func NewBarangService(repo IBarangRepository) *barangService {
	return &barangService{repo: repo}
}

func (s *barangService) GetAll() ([]models.Barang, error) {
	return s.repo.GetBarang()
}

func (s *barangService) GetByID(id uint) (*models.Barang, error) {
	return s.repo.FindByID(id)
}

func (s *barangService) Create(data *models.Barang) error {
	data.KodeBarang = utils.GenerateID(config.DB, "MTR", true)
	return s.repo.CreateBarang(*data)
}

func (s *barangService) Update(id uint, data *models.Barang) error {
	barang, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	data.KodeBarang = barang.KodeBarang
	data.ID = barang.ID
	return s.repo.Update(data)
}

func (s *barangService) Delete(id uint) error {
	return s.repo.Delete(id)
}
