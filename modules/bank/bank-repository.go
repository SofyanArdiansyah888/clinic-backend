package bank

import (
	"backend/models"
	"gorm.io/gorm"
)

type IBankRepository interface {
	FindAll() ([]models.Bank, error)
	FindByID(id uint) (*models.Bank, error)
	Create(bank *models.Bank) error
	Update(bank *models.Bank) error
	Delete(id uint) error
}

type bankRepository struct {
	db *gorm.DB
}

func NewBankRepository(db *gorm.DB) IBankRepository {
	return &bankRepository{db: db}
}

func (r *bankRepository) FindAll() ([]models.Bank, error) {
	var banks []models.Bank
	err := r.db.Find(&banks).Error
	return banks, err
}

func (r *bankRepository) FindByID(id uint) (*models.Bank, error) {
	var bank models.Bank
	err := r.db.First(&bank, id).Error
	return &bank, err
}

func (r *bankRepository) Create(bank *models.Bank) error {
	return r.db.Create(bank).Error
}

func (r *bankRepository) Update(bank *models.Bank) error {
	return r.db.Model(&models.Bank{}).Where("id = ?", bank.ID).Updates(bank).Error
}

func (r *bankRepository) Delete(id uint) error {
	return r.db.Delete(&models.Bank{}, id).Error
}
