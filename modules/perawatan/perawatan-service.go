package perawatan

import (
	"backend/models"
)

type IPerawatanService interface {
	GetAll() ([]models.Perawatan, error)
	GetByID(id uint) (*models.Perawatan, error)
	Create(data *models.Perawatan) error
	Update(id uint, data *models.Perawatan) error
	Delete(id uint) error
}

type perawatanService struct {
	repo IPerawatanRepository
}

func NewPerawatanService(repo IPerawatanRepository) *perawatanService {
	return &perawatanService{repo: repo}
}

func (s *perawatanService) GetAll() ([]models.Perawatan, error) {
	return s.repo.FindAll()
}

func (s *perawatanService) GetByID(id uint) (*models.Perawatan, error) {
	return s.repo.FindByID(id)
}

func (s *perawatanService) Create(data *models.Perawatan) error {
	//data.NoRM = utils.GenerateID(config.DB, "RMD", true)
	//data.NoMember = utils.GenerateID(config.DB, "MEM", true)
	return s.repo.Create(data)
}

func (s *perawatanService) Update(id uint, data *models.Perawatan) error {
	staff, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	//staff.Nama = data.Nama
	//staff.Email = data.Email
	//staff.Kota = data.Kota
	return s.repo.Update(staff)
}

func (s *perawatanService) Delete(id uint) error {
	return s.repo.Delete(id)
}
