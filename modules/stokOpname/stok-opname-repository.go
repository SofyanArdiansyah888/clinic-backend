package stokOpname

import (
	"backend/models"

	"gorm.io/gorm"
)

type IStokOpnameRepository interface {
	FindAll() ([]models.StokOpname, error)
	FindByID(id uint) (*models.StokOpname, error)
	Create(stokOpname *models.StokOpname) error
}

type stokOpnameRepository struct {
	db *gorm.DB
}

func NewStokOpnameRepository(db *gorm.DB) IStokOpnameRepository {
	return &stokOpnameRepository{db: db}
}

func (r *stokOpnameRepository) FindAll() ([]models.StokOpname, error) {
	var stokOpname []models.StokOpname
	err := r.db.Find(&stokOpname).Error
	return stokOpname, err
}

func (r *stokOpnameRepository) FindByID(id uint) (*models.StokOpname, error) {
	var stokOpname models.StokOpname
	err := r.db.First(&stokOpname, id).Error
	return &stokOpname, err
}

func (r *stokOpnameRepository) Create(stokOpname *models.StokOpname) error {
	return r.db.Create(stokOpname).Error
}
