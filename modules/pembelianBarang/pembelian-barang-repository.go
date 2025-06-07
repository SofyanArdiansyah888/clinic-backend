package pembelianBarang

import (
	"backend/models"

	"gorm.io/gorm"
)

type PembelianBarangRepository struct {
	db *gorm.DB
}

func NewPembelianBarangRepository(db *gorm.DB) *PembelianBarangRepository {
	return &PembelianBarangRepository{db: db}
}

// Repository: Fokus pada operasi database
// Add this method to PembelianBarangRepository
func (r *PembelianBarangRepository) CreateTransaksiAndUpdateStock(transaksi *models.Pembelian, details []models.PembelianDetail) error {
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

		return nil
	})
}

func (r *PembelianBarangRepository) FindByNomor(nomorTransaksi string) (*models.Pembelian, []models.PembelianDetail, error) {
	var transaksi models.Pembelian
	var details []models.PembelianDetail

	// Get header
	if err := r.db.Where("no_transaksi = ?", nomorTransaksi).First(&transaksi).Error; err != nil {
		return nil, nil, err
	}

	// Get details
	if err := r.db.Where("no_transaksi = ?", nomorTransaksi).Find(&details).Error; err != nil {
		return nil, nil, err
	}

	// Update stock movement for each purchased item
	for _, detail := range details {
		movement := models.StokMovement{
			KodeBarang:    detail.KodeBarang,
			KodeReferensi: transaksi.NoTransaksi,
			Quantity:      detail.Jumlah,
			Jenis:         "pembelian",
			Keterangan:    "Pembelian dari supllier",
		}

		if err := r.db.Create(&movement).Error; err != nil {
			return nil, nil, err
		}
	}

	return &transaksi, details, nil
}
