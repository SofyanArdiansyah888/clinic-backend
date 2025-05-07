package perawatan

import (
	"backend/models"
	"gorm.io/gorm"
)

type IPerawatanRepository interface {
	FindAll() ([]models.Perawatan, error)
	FindByID(id uint) (*models.Perawatan, error)
	Create(pasien *models.Perawatan) error
	Update(pasien *models.Perawatan) error
	Delete(id uint) error
}

type pasienRepository struct {
	db *gorm.DB
}

func NewPerawatanRepository(db *gorm.DB) IPerawatanRepository {
	return &pasienRepository{db: db}
}

func (r *pasienRepository) FindAll() ([]models.Perawatan, error) {
	var pasiens []models.Perawatan
	err := r.db.Find(&pasiens).Error
	return pasiens, err
}

func (r *pasienRepository) FindByID(id uint) (*models.Perawatan, error) {
	var pasien models.Perawatan
	err := r.db.First(&pasien, id).Error
	return &pasien, err
}

func (r *pasienRepository) Create(pasien *models.Perawatan) error {
	return r.db.Create(pasien).Error
}

func (r *pasienRepository) Update(pasien *models.Perawatan) error {
	return r.db.Save(pasien).Error
}

func (r *pasienRepository) Delete(id uint) error {
	return r.db.Delete(&models.Perawatan{}, id).Error
}
