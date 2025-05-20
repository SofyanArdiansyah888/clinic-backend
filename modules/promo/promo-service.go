package promo

import "backend/models"

type PromoService struct {
	repository *PromoRepository
}

func NewPromoService(repository *PromoRepository) *PromoService {
	return &PromoService{repository}
}

func (s *PromoService) GetAll() ([]models.Promo, error) {
	return s.repository.FindAll()
}

func (s *PromoService) GetByID(id uint) (models.Promo, error) {
	return s.repository.FindByID(id)
}

func (s *PromoService) Create(promo *models.Promo) error {
	return s.repository.Create(promo)
}

func (s *PromoService) Update(promo *models.Promo) error {
	return s.repository.Update(promo)
}

func (s *PromoService) Delete(id uint) error {
	return s.repository.Delete(id)
}
