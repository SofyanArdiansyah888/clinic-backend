package perusahaan

import (
	"backend/models"
)

type IPerusahaanService interface {
	GetByID(id uint) (*models.Perusahaan, error)
	Create(data *models.Perusahaan) error
	Update(id uint, data *models.Perusahaan) error
}

type perusahaanService struct {
	repo IPerusahaanRepository
}

func NewPerusahaanService(repo IPerusahaanRepository) *perusahaanService {
	return &perusahaanService{repo: repo}
}

func (s *perusahaanService) GetByID(id uint) (*models.Perusahaan, error) {
	return s.repo.FindByID(id)
}

func (s *perusahaanService) Create(data *models.Perusahaan) error {
	//data.NoRM = utils.GenerateID(config.DB, "RMD", true)
	//data.NoMember = utils.GenerateID(config.DB, "MEM", true)
	return s.repo.Create(data)
}

func (s *perusahaanService) Update(id uint, data *models.Perusahaan) error {
	staff, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	//staff.Nama = data.Nama
	//staff.Email = data.Email
	//staff.Kota = data.Kota
	return s.repo.Update(staff)
}
