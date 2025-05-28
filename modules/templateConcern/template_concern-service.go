package templateConcern

import "backend/models"

type TemplateConcernService struct {
	repository *TemplateConcernRepository
}

func NewTemplateConcernService(repository *TemplateConcernRepository) *TemplateConcernService {
	return &TemplateConcernService{repository}
}

func (s *TemplateConcernService) GetAll() ([]models.TemplateConcern, error) {
	return s.repository.FindAll()
}

func (s *TemplateConcernService) GetByID(id uint) (models.TemplateConcern, error) {
	return s.repository.FindByID(id)
}

func (s *TemplateConcernService) Create(concern *models.TemplateConcern) error {
	return s.repository.Create(concern)
}

func (s *TemplateConcernService) Update(concern *models.TemplateConcern) error {
	return s.repository.Update(concern)
}

func (s *TemplateConcernService) Delete(id uint) error {
	return s.repository.Delete(id)
}
