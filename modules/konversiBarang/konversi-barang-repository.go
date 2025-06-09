package konversiBarang

import (
	"backend/models"

	"gorm.io/gorm"
)

type KonversiBarangRepository struct {
	db *gorm.DB
}

func NewKonversiBarangRepository(db *gorm.DB) *KonversiBarangRepository {
	return &KonversiBarangRepository{db: db}
}

func (r *KonversiBarangRepository) FindAll() ([]models.KonversiStok, error) {
	var konversis []models.KonversiStok
	err := r.db.Find(&konversis).Error
	return konversis, err
}

// Repository: Fokus pada operasi database
func (r *KonversiBarangRepository) CreateTransaksiAndUpdateStock(transaksi *models.KonversiStok, details []models.KonversiStokDetail) error {
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

		// Update stock movement for converted items
		for _, detail := range details {
			if detail.Arah == "keluar" {
				// Decrease stock for source item
				sourceMovement := models.StokMovement{
					KodeBarang:    detail.KodeBarang,
					KodeReferensi: transaksi.NoKonversi,
					Quantity:      -int(detail.Quantity), // negative for reduction
					Jenis:         "konversi keluar",
					Keterangan:    "Konversi keluar dengan nomor referensi - " + transaksi.NoKonversi,
				}

				if err := tx.Create(&sourceMovement).Error; err != nil {
					return err
				}
			}
			if detail.Arah == "masuk" {
				// Increase stock for target item
				targetMovement := models.StokMovement{
					KodeBarang:    detail.KodeBarang,
					KodeReferensi: transaksi.NoKonversi,
					Quantity:      int(detail.Quantity),
					Jenis:         "konversi masuk",
					Keterangan:    "Konversi masuk dengan nomor referensi - " + transaksi.NoKonversi,
				}

				if err := tx.Create(&targetMovement).Error; err != nil {
					return err
				}
			}

		}

		return nil
	})
}

func (r *KonversiBarangRepository) FindByNomor(nomorTransaksi string) (*models.KonversiStok, []models.KonversiStokDetail, error) {
	var transaksi models.KonversiStok
	var details []models.KonversiStokDetail

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
