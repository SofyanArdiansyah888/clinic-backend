package kartuStok

import (
	"backend/models"

	"gorm.io/gorm"
)

type IKartuStokRepository interface {
	GetStokMovement(kodeKartuStok string, dari string, sampai string) ([]models.StokMovement, error)
}
type kartuStokRepository struct {
	db *gorm.DB
}

func NewKartuStokRepository(db *gorm.DB) IKartuStokRepository {
	return &kartuStokRepository{db: db}
}

func (repo *kartuStokRepository) GetStokMovement(kodeKartuStok string, dari string, sampai string) ([]models.StokMovement, error) {
	var stokMovements []models.StokMovement
	result := repo.db.
		Where("kode_barang = ?", kodeKartuStok).
		Where("created_at BETWEEN ? AND ?", dari, sampai).
		Order("created_at desc").
		Find(&stokMovements)
	if result.Error != nil {
		return []models.StokMovement{}, result.Error
	}
	if len(stokMovements) == 0 {
		return []models.StokMovement{}, nil
	}
	return stokMovements, nil
}
