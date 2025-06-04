package stokOpname

import (
	"backend/models"

	"gorm.io/gorm"
)

type IStokOpnameRepository interface {
	FindAll() ([]models.StokOpname, error)
	FindByID(id uint) (*models.StokOpname, error)
	CreateAndUpdateStok(stokOpname *models.StokOpname) error
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

func (r *stokOpnameRepository) CreateAndUpdateStok(stokOpname *models.StokOpname) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		// Create new stock opname record
		if err := tx.Create(stokOpname).Error; err != nil {
			return err
		}

		// Update actual stock quantity in barang table
		if err := tx.Model(&models.Barang{}).
			Where("kode_barang = ?", stokOpname.Barang.KodeBarang).
			Update("stok", stokOpname.StokRiil).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}
