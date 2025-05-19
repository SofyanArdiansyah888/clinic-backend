package cabang

import (
	"backend/config"
	"backend/models"
	"backend/utils"
)

type ICabangService interface {
	GetAll() ([]models.Cabang, error)
	GetByID(id uint) (*models.Cabang, error)
	Create(data *models.Cabang) error
	Update(id uint, data *models.Cabang) error
	Delete(id uint) error
}

type cabangService struct {
	repo ICabangRepository
}

func NewCabangService(repo ICabangRepository) *cabangService {
	return &cabangService{repo: repo}
}

func (s *cabangService) GetAll() ([]models.Cabang, error) {
	return s.repo.FindAll()
}

func (s *cabangService) GetByID(id uint) (*models.Cabang, error) {
	return s.repo.FindByID(id)
}

func (s *cabangService) Create(data *models.Cabang) error {
	data.NoCabang = utils.GenerateID(config.DB, "CBG", true)
	return s.repo.Create(data)
}

func (s *cabangService) Update(id uint, data *models.Cabang) error {
	cabang, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	data.NoCabang = cabang.NoCabang
	return s.repo.Update(data)
}

func (s *cabangService) Delete(id uint) error {
	return s.repo.Delete(id)
}