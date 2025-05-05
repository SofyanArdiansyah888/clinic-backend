package riwayatPerawatan

import (
	"backend/models"
	"gorm.io/gorm"
)

type IRiwayatPerawatanRepository interface {
	FindAll() ([]models.RiwayatPerawatan, error)
	FindByID(id uint) (*models.RiwayatPerawatan, error)
	Create(pasien *models.RiwayatPerawatan) error
	Update(pasien *models.RiwayatPerawatan) error
	Delete(id uint) error
}

type pasienRepository struct {
	db *gorm.DB
}

func NewRiwayatPerawatanRepository(db *gorm.DB) IRiwayatPerawatanRepository {
	return &pasienRepository{db: db}
}

func (r *pasienRepository) FindAll() ([]models.RiwayatPerawatan, error) {
	var pasiens []models.RiwayatPerawatan
	err := r.db.Find(&pasiens).Error
	return pasiens, err
}

func (r *pasienRepository) FindByID(id uint) (*models.RiwayatPerawatan, error) {
	var pasien models.RiwayatPerawatan
	err := r.db.First(&pasien, id).Error
	return &pasien, err
}

func (r *pasienRepository) Create(pasien *models.RiwayatPerawatan) error {
	return r.db.Create(pasien).Error
}

func (r *pasienRepository) Update(pasien *models.RiwayatPerawatan) error {
	return r.db.Save(pasien).Error
}

func (r *pasienRepository) Delete(id uint) error {
	return r.db.Delete(&models.RiwayatPerawatan{}, id).Error
}
