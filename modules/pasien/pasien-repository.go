package pasien

import (
	"backend/models"
	"gorm.io/gorm"
)

type IPasienRepository interface {
	FindAll() ([]models.Pasien, error)
	FindByID(id uint) (*models.Pasien, error)
	Create(pasien *models.Pasien) error
	Update(pasien *models.Pasien) error
	Delete(id uint) error
}

type pasienRepository struct {
	db *gorm.DB
}

func NewPasienRepository(db *gorm.DB) IPasienRepository {
	return &pasienRepository{db: db}
}

func (r *pasienRepository) FindAll() ([]models.Pasien, error) {
	var pasiens []models.Pasien
	err := r.db.Find(&pasiens).Error
	return pasiens, err
}

func (r *pasienRepository) FindByID(id uint) (*models.Pasien, error) {
	var pasien models.Pasien
	err := r.db.First(&pasien, id).Error
	return &pasien, err
}

func (r *pasienRepository) Create(pasien *models.Pasien) error {
	return r.db.Create(pasien).Error
}

func (r *pasienRepository) Update(pasien *models.Pasien) error {
	return r.db.Save(pasien).Error
}

func (r *pasienRepository) Delete(id uint) error {
	return r.db.Delete(&models.Pasien{}, id).Error
}
