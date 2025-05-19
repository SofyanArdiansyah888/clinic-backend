package cabang

import (
	"backend/models"
	"gorm.io/gorm"
)

type ICabangRepository interface {
	FindAll() ([]models.Cabang, error)
	FindByID(id uint) (*models.Cabang, error)
	Create(cabang *models.Cabang) error
	Update(cabang *models.Cabang) error
	Delete(id uint) error
}

type cabangRepository struct {
	db *gorm.DB
}

func NewCabangRepository(db *gorm.DB) ICabangRepository {
	return &cabangRepository{db: db}
}

func (r *cabangRepository) FindAll() ([]models.Cabang, error) {
	var cabang []models.Cabang
	err := r.db.Find(&cabang).Error
	return cabang, err
}

func (r *cabangRepository) FindByID(id uint) (*models.Cabang, error) {
	var cabang models.Cabang
	err := r.db.First(&cabang, id).Error
	return &cabang, err
}

func (r *cabangRepository) Create(cabang *models.Cabang) error {
	return r.db.Create(cabang).Error
}

func (r *cabangRepository) Update(cabang *models.Cabang) error {
	return r.db.Model(&models.Cabang{}).Where("id = ?", cabang.ID).Updates(cabang).Error
}

func (r *cabangRepository) Delete(id uint) error {
	return r.db.Delete(&models.Cabang{}, id).Error
}