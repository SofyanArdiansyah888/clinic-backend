package treatment

import (
	"backend/models"
	"gorm.io/gorm"
)

type ITreatmentRepository interface {
	FindAll() ([]models.Treatment, error)
	FindByID(id uint) (*models.Treatment, error)
	Create(treatment *models.Treatment) error
	Update(treatment *models.Treatment) error
	Delete(id uint) error
}

type treatmentRepository struct {
	db *gorm.DB
}

func NewTreatmentRepository(db *gorm.DB) ITreatmentRepository {
	return &treatmentRepository{db: db}
}

func (r *treatmentRepository) FindAll() ([]models.Treatment, error) {
	var treatment []models.Treatment
	err := r.db.Find(&treatment).Error
	return treatment, err
}

func (r *treatmentRepository) FindByID(id uint) (*models.Treatment, error) {
	var treatment models.Treatment
	err := r.db.First(&treatment, id).Error
	return &treatment, err
}

func (r *treatmentRepository) Create(treatment *models.Treatment) error {
	return r.db.Create(treatment).Error
}

func (r *treatmentRepository) Update(treatment *models.Treatment) error {
	return r.db.Model(&models.Treatment{}).Where("id = ?", treatment.ID).Updates(treatment).Error
}

func (r *treatmentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Treatment{}, id).Error
}
