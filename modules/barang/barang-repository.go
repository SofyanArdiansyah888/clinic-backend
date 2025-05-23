package barang

import (
	"backend/models"

	"gorm.io/gorm"
)

type IBarangRepository interface {
	GetBarang() ([]models.Barang, error)
	FindByID(id uint) (*models.Barang, error)
	CreateBarang(barang models.Barang) error
	Update(barang *models.Barang) error
	Delete(id uint) error
}
type barangRepository struct {
	db *gorm.DB
}

func NewBarangRepository(db *gorm.DB) IBarangRepository {
	return &barangRepository{db: db}
}

func (b *barangRepository) GetBarang() ([]models.Barang, error) {
	var barang []models.Barang
	err := b.db.Find(&barang).Error
	return barang, err
}

func (b *barangRepository) FindByID(id uint) (*models.Barang, error) {
	var barang models.Barang
	err := b.db.First(&barang, id).Error
	return &barang, err
}

func (b *barangRepository) CreateBarang(barang models.Barang) error {
	return b.db.Create(&barang).Error
}

func (b *barangRepository) Update(barang *models.Barang) error {
	return b.db.Model(&models.Barang{}).Where("id = ?", barang.ID).Updates(barang).Error
}

func (b *barangRepository) Delete(id uint) error {
	return b.db.Delete(&models.Barang{}, id).Error
}
