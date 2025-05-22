package templateConcern

import (
	"backend/models"
	"gorm.io/gorm"
)

type TemplateConcernRepository struct {
	db *gorm.DB
}

func NewTemplateConcernRepository(db *gorm.DB) *TemplateConcernRepository {
	return &TemplateConcernRepository{db}
}

func (r *TemplateConcernRepository) FindAll() ([]models.TemplateConcern, error) {
	var concerns []models.TemplateConcern
	err := r.db.Find(&concerns).Error
	return concerns, err
}

func (r *TemplateConcernRepository) FindByID(id uint) (models.TemplateConcern, error) {
	var concern models.TemplateConcern
	err := r.db.First(&concern, id).Error
	return concern, err
}

func (r *TemplateConcernRepository) Create(concern *models.TemplateConcern) error {
	return r.db.Create(concern).Error
}

func (r *TemplateConcernRepository) Update(concern *models.TemplateConcern) error {
	return r.db.Save(concern).Error
}

func (r *TemplateConcernRepository) Delete(id uint) error {
	return r.db.Delete(&models.TemplateConcern{}, id).Error
}