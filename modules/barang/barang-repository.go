package barang

import (
	"backend/config"
	"backend/models"
)

type IBarangRepository interface {
	GetBarang() ([]models.Barang, error)
	CreateBarang(barang models.Barang) (models.Barang, error)
}
type barangRepository struct{}

func NewBarangRepository() IBarangRepository {
	return &barangRepository{}
}

func (b *barangRepository) GetBarang() ([]models.Barang, error) {
	var barang []models.Barang
	err := config.DB.Find(&barang).Error
	return barang, err
}

func (b *barangRepository) CreateBarang(barang models.Barang) (models.Barang, error) {
	//TODO implement me
	panic("implement me")
}
