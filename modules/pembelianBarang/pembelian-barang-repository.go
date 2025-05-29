package pembelianBarang

import (
	"gorm.io/gorm"
	"backend/models"
)

type PembelianBarangRepository struct {
	db *gorm.DB
}

func NewPembelianBarangRepository(db *gorm.DB) *PembelianBarangRepository {
	return &PembelianBarangRepository{db: db}
}

func (r *PembelianBarangRepository) Create(transaksi *models.TransaksiBarang, details []models.TransaksiBarangDetail) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Simpan header transaksi
		if err := tx.Create(transaksi).Error; err != nil {
			return err
		}

		// Simpan detail transaksi
		for _, detail := range details {
			detail.NomorTransaksi = transaksi.NomorTransaksi
			detail.IDCabang = transaksi.IDCabang
			detail.Tipe = transaksi.Tipe
			if err := tx.Create(&detail).Error; err != nil {
				return err
			}
		}

		return nil
	})
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