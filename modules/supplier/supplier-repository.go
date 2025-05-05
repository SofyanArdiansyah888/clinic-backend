package supplier

import (
	"backend/models"
	"gorm.io/gorm"
)

type ISupplierRepository interface {
	FindAll() ([]models.Supplier, error)
	FindByID(id uint) (*models.Supplier, error)
	Create(supplier *models.Supplier) error
	Update(supplier *models.Supplier) error
	Delete(id uint) error
}

type supplierRepository struct {
	db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) ISupplierRepository {
	return &supplierRepository{db: db}
}

func (r *supplierRepository) FindAll() ([]models.Supplier, error) {
	var suppliers []models.Supplier
	err := r.db.Find(&suppliers).Error
	return suppliers, err
}

func (r *supplierRepository) FindByID(id uint) (*models.Supplier, error) {
	var supplier models.Supplier
	err := r.db.First(&supplier, id).Error
	return &supplier, err
}

func (r *supplierRepository) Create(supplier *models.Supplier) error {
	return r.db.Create(supplier).Error
}

func (r *supplierRepository) Update(supplier *models.Supplier) error {
	return r.db.Model(&models.Supplier{}).Where("id = ?", supplier.ID).Updates(supplier).Error
}

func (r *supplierRepository) Delete(id uint) error {
	return r.db.Delete(&models.Supplier{}, id).Error
}
