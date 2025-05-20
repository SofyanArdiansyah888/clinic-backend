package promo

import (
	"backend/models"
	"gorm.io/gorm"
)

type PromoRepository struct {
	db *gorm.DB
}

func NewPromoRepository(db *gorm.DB) *PromoRepository {
	return &PromoRepository{db}
}

func (r *PromoRepository) FindAll() ([]models.Promo, error) {
	var promos []models.Promo
	err := r.db.Find(&promos).Error
	return promos, err
}

func (r *PromoRepository) FindByID(id uint) (models.Promo, error) {
	var promo models.Promo
	err := r.db.First(&promo, id).Error
	return promo, err
}

func (r *PromoRepository) Create(promo *models.Promo) error {
	return r.db.Create(promo).Error
}

func (r *PromoRepository) Update(promo *models.Promo) error {
	return r.db.Save(promo).Error
}

func (r *PromoRepository) Delete(id uint) error {
	return r.db.Delete(&models.Promo{}, id).Error
}
