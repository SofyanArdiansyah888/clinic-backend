package pasien

import (
	"backend/config"
	"backend/models"
	"backend/utils"
)

type IPasienService interface {
	GetAll() ([]models.Pasien, error)
	GetByID(id uint) (*models.Pasien, error)
	Create(data *models.Pasien) error
	Update(id uint, data *models.Pasien) error
	Delete(id uint) error
}

type pasienService struct {
	repo IPasienRepository
}

func NewPasienService(repo IPasienRepository) *pasienService {
	return &pasienService{repo: repo}
}

func (s *pasienService) GetAll() ([]models.Pasien, error) {
	return s.repo.FindAll()
}

func (s *pasienService) GetByID(id uint) (*models.Pasien, error) {
	return s.repo.FindByID(id)
}

func (s *pasienService) Create(data *models.Pasien) error {
	data.NoRM = utils.GenerateID(config.DB, "RMD", true)
	data.NoMember = utils.GenerateID(config.DB, "MEM", true)
	return s.repo.Create(data)
}

func (s *pasienService) Update(id uint, data *models.Pasien) error {
	staff, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	//staff.Nama = data.Nama
	//staff.Email = data.Email
	//staff.Kota = data.Kota
	return s.repo.Update(staff)
}

func (s *pasienService) Delete(id uint) error {
	return s.repo.Delete(id)
}
