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
func (r *PembelianBarangRepository) CreateTransaksiAndUpdateStock(transaksi *models.TransaksiBarang, details []models.TransaksiBarangDetail, stockUpdates map[string]int) error {
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

		// Update stock
		for kodeBarang, newStock := range stockUpdates {
			if err := tx.Model(&models.Barang{}).Where("kode_barang = ?", kodeBarang).Update("stok", newStock).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *PembelianBarangRepository) GetStokByKodeBarang(kodeBarang string) (int, error) {
	var barang models.Barang
	err := r.db.Select("stok").Where("kode_barang = ?", kodeBarang).First(&barang).Error
	if err != nil {
		return 0, err
	}
	return barang.Stok, nil
}

func (r *PembelianBarangRepository) FindByNomor(nomorTransaksi string) (*models.TransaksiBarang, []models.TransaksiBarangDetail, error) {
	var transaksi models.TransaksiBarang
	var details []models.TransaksiBarangDetail

	// Get header
	if err := r.db.Where("nomor_transaksi = ?", nomorTransaksi).First(&transaksi).Error; err != nil {
		return nil, nil, err
	}

	// Get details
	if err := r.db.Where("nomor_transaksi = ?", nomorTransaksi).Find(&details).Error; err != nil {
		return nil, nil, err
	}

	return &transaksi, details, nil
}
