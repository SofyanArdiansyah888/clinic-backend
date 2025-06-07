package produksiBarang

import (
	"backend/models"

	"gorm.io/gorm"
)

type ProduksiBarangRepository struct {
	db *gorm.DB
}

func NewProduksiBarangRepository(db *gorm.DB) *ProduksiBarangRepository {
	return &ProduksiBarangRepository{db: db}
}

// Repository: Fokus pada operasi database
func (r *ProduksiBarangRepository) CreateTransaksiAndUpdateStock(transaksi *models.ProduksiBarang, details []models.ProduksiBarangDetail) error {
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

		// Update stock movement for produced items
		for _, detail := range details {
			// Decrease stock for raw materials
			sourceMovement := models.StokMovement{
				KodeBarang:    detail.KodeBarang,
				KodeReferensi: transaksi.NoProduksi,
				Quantity:      -int(detail.Quantity), // negative for reduction
				Jenis:         "produksi keluar",
				Keterangan:    "Bahan baku keluar untuk produksi dengan nomor - " + transaksi.NoProduksi,
			}

			if err := tx.Create(&sourceMovement).Error; err != nil {
				return err
			}

			// Increase stock for finished product
			targetMovement := models.StokMovement{
				KodeBarang:    detail.KodeBarang,
				KodeReferensi: transaksi.NoProduksi,
				Quantity:      int(detail.Quantity),
				Jenis:         "produksi masuk",
				Keterangan:    "Hasil produksi masuk dengan nomor - " + transaksi.NoProduksi,
			}

			if err := tx.Create(&targetMovement).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *ProduksiBarangRepository) FindByNomor(nomorTransaksi string) (*models.ProduksiBarang, []models.ProduksiBarangDetail, error) {
	var transaksi models.ProduksiBarang
	var details []models.ProduksiBarangDetail

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
