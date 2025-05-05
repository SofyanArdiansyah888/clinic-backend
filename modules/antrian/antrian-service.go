package antrian

import (
	"backend/models"
)

type IAntrianService interface {
	GetAll() ([]models.Antrian, error)
	GetByID(id uint) (*models.Antrian, error)
	Create(data *models.Antrian) error
	Update(id uint, data *models.Antrian) error
	Delete(id uint) error
}

type antrianService struct {
	repo IAntrianRepository
}

func NewAntrianService(repo IAntrianRepository) *antrianService {
	return &antrianService{repo: repo}
}

func (s *antrianService) GetAll() ([]models.Antrian, error) {
	return s.repo.FindAll()
}

func (s *antrianService) GetByID(id uint) (*models.Antrian, error) {
	return s.repo.FindByID(id)
}

func (s *antrianService) Create(data *models.Antrian) error {
	return s.repo.Create(data)
}

func (s *antrianService) Update(id uint, data *models.Antrian) error {
	antrian, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	data.NoAntrian = antrian.NoAntrian
	data.ID = antrian.ID
	return s.repo.Update(data)
}

func (s *antrianService) Delete(id uint) error {
	return s.repo.Delete(id)
}
