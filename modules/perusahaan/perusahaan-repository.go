package perusahaan

import (
	"backend/models"
	"gorm.io/gorm"
)

type IPerusahaanRepository interface {
	FindByID(id uint) (*models.Perusahaan, error)
	Create(perusahaan *models.Perusahaan) error
	Update(perusahaan *models.Perusahaan) error
}

type perusahaanRepository struct {
	db *gorm.DB
}

func NewPerusahaanRepository(db *gorm.DB) IPerusahaanRepository {
	return &perusahaanRepository{db: db}
}

func (r *perusahaanRepository) FindByID(id uint) (*models.Perusahaan, error) {
	var perusahaan models.Perusahaan
	err := r.db.First(&perusahaan, id).Error
	return &perusahaan, err
}

func (r *perusahaanRepository) Create(perusahaan *models.Perusahaan) error {
	return r.db.Create(perusahaan).Error
}

func (r *perusahaanRepository) Update(perusahaan *models.Perusahaan) error {
	return r.db.Save(perusahaan).Error
}
