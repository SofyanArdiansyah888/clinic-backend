package treatment

import (
	"backend/config"
	"backend/models"
	"backend/utils"
)

type ITreatmentService interface {
	GetAll() ([]models.Treatment, error)
	GetByID(id uint) (*models.Treatment, error)
	Create(data *models.Treatment) error
	Update(id uint, data *models.Treatment) error
	Delete(id uint) error
}

type treatmentService struct {
	repo ITreatmentRepository
}

func NewTreatmentService(repo ITreatmentRepository) *treatmentService {
	return &treatmentService{repo: repo}
}

func (s *treatmentService) GetAll() ([]models.Treatment, error) {
	return s.repo.FindAll()
}

func (s *treatmentService) GetByID(id uint) (*models.Treatment, error) {
	return s.repo.FindByID(id)
}

func (s *treatmentService) Create(data *models.Treatment) error {
	data.NoTreatment = utils.GenerateID(config.DB, "TRT", true)
	return s.repo.Create(data)
}

func (s *treatmentService) Update(id uint, data *models.Treatment) error {
	treatment, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	data.NoTreatment = treatment.NoTreatment
	return s.repo.Update(data)
}

func (s *treatmentService) Delete(id uint) error {
	return s.repo.Delete(id)
}
