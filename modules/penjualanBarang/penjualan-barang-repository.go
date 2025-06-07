package penjualanBarang

import (
	"backend/models"

	"gorm.io/gorm"
)

type PenjualanBarangRepository struct {
	db *gorm.DB
}

func NewPenjualanBarangRepository(db *gorm.DB) *PenjualanBarangRepository {
	return &PenjualanBarangRepository{db: db}
}

// Repository: Fokus pada operasi database
// Add this method to PenjualanBarangRepository
func (r *PenjualanBarangRepository) CreateTransaksiAndUpdateStock(transaksi *models.Penjualan, details []models.PenjualanDetail) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Create transaksi
		if err := tx.Create(transaksi).Error; err != nil {
			return err
		}

		// Create details
		for _, detail := range details {
			if err := tx.Create(&detail).Error; err != nil {
				return err
			}
		}

		// Update stock movement for each purchased item
		for _, detail := range details {
			movement := models.StokMovement{
				KodeBarang:    detail.KodeBarang,
				KodeReferensi: transaksi.NoTransaksi,
				Quantity:      -detail.Jumlah, // Negative quantity for sales
				Jenis:         "penjualan",
				Keterangan:    "Penjualan barang dengan nomor referensi - " + transaksi.NoTransaksi,
			}

			if err := r.db.Create(&movement).Error; err != nil {
				return nil
			}
		}

		return nil
	})
}

func (r *PenjualanBarangRepository) FindByNomor(nomorTransaksi string) (*models.Penjualan, []models.PenjualanDetail, error) {
	var transaksi models.Penjualan
	var details []models.PenjualanDetail

	// Get header
	if err := r.db.Where("no_transaksi = ?", nomorTransaksi).First(&transaksi).Error; err != nil {
		return nil, nil, err
	}

	// Get details
	if err := r.db.Where("no_transaksi = ?", nomorTransaksi).Find(&details).Error; err != nil {
		return nil, nil, err
	}

	return &transaksi, details, nil
}
