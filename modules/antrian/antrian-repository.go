package antrian

import (
	"backend/models"
	"gorm.io/gorm"
)

type IAntrianRepository interface {
	FindAll() ([]models.Antrian, error)
	FindByID(id uint) (*models.Antrian, error)
	Create(antrian *models.Antrian) error
	Update(antrian *models.Antrian) error
	Delete(id uint) error
}

type antrianRepository struct {
	db *gorm.DB
}

func NewAntrianRepository(db *gorm.DB) IAntrianRepository {
	return &antrianRepository{db: db}
}

func (r *antrianRepository) FindAll() ([]models.Antrian, error) {
	var antrian []models.Antrian
	err := r.db.Find(&antrian).Error
	return antrian, err
}

func (r *antrianRepository) FindByID(id uint) (*models.Antrian, error) {
	var antrian models.Antrian
	err := r.db.First(&antrian, id).Error
	return &antrian, err
}

func (r *antrianRepository) Create(antrian *models.Antrian) error {
	return r.db.Omit("Pasien", "Staff").Create(antrian).Error
}

func (r *antrianRepository) Update(antrian *models.Antrian) error {
	return r.db.Model(&models.Antrian{}).Where("id = ?", antrian.ID).Updates(antrian).Error
}

func (r *antrianRepository) Delete(id uint) error {
	return r.db.Delete(&models.Antrian{}, id).Error
}
