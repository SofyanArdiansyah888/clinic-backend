package riwayatPerawatan

import (
	"backend/models"
)

type IRiwayatPerawatanService interface {
	GetAll() ([]models.RiwayatPerawatan, error)
	GetByID(id uint) (*models.RiwayatPerawatan, error)
	Create(data *models.RiwayatPerawatan) error
	Update(id uint, data *models.RiwayatPerawatan) error
	Delete(id uint) error
}

type riwayatPerawatanService struct {
	repo IRiwayatPerawatanRepository
}

func NewRiwayatPerawatanService(repo IRiwayatPerawatanRepository) *riwayatPerawatanService {
	return &riwayatPerawatanService{repo: repo}
}

func (s *riwayatPerawatanService) GetAll() ([]models.RiwayatPerawatan, error) {
	return s.repo.FindAll()
}

func (s *riwayatPerawatanService) GetByID(id uint) (*models.RiwayatPerawatan, error) {
	return s.repo.FindByID(id)
}

func (s *riwayatPerawatanService) Create(data *models.RiwayatPerawatan) error {
	//data.NoRM = utils.GenerateID(config.DB, "RMD", true)
	//data.NoMember = utils.GenerateID(config.DB, "MEM", true)
	return s.repo.Create(data)
}

func (s *riwayatPerawatanService) Update(id uint, data *models.RiwayatPerawatan) error {
	staff, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	//staff.Nama = data.Nama
	//staff.Email = data.Email
	//staff.Kota = data.Kota
	return s.repo.Update(staff)
}

func (s *riwayatPerawatanService) Delete(id uint) error {
	return s.repo.Delete(id)
}
