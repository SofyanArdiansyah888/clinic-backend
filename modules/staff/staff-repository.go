package staff

import (
	"backend/models"
	"gorm.io/gorm"
)

type IStaffRepository interface {
	FindAll() ([]models.Staff, error)
	FindByID(id uint) (*models.Staff, error)
	Create(staff *models.Staff) error
	Update(staff *models.Staff) error
	Delete(id uint) error
}

type staffRepository struct {
	db *gorm.DB
}

func NewStaffRepository(db *gorm.DB) IStaffRepository {
	return &staffRepository{db: db}
}

func (r *staffRepository) FindAll() ([]models.Staff, error) {
	var staffs []models.Staff
	err := r.db.Find(&staffs).Error
	return staffs, err
}

func (r *staffRepository) FindByID(id uint) (*models.Staff, error) {
	var staff models.Staff
	err := r.db.First(&staff, id).Error
	return &staff, err
}

func (r *staffRepository) Create(staff *models.Staff) error {
	return r.db.Create(staff).Error
}

func (r *staffRepository) Update(staff *models.Staff) error {
	return r.db.Model(&models.Staff{}).Where("id = ?", staff.ID).Updates(staff).Error
}

func (r *staffRepository) Delete(id uint) error {
	return r.db.Delete(&models.Staff{}, id).Error
}
