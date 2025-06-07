package konversiBarang

import (
	"backend/config"
	"backend/models"
	"backend/utils"
)

type KonversiBarangService struct {
	repo *KonversiBarangRepository
}

func NewKonversiBarangService(repo *KonversiBarangRepository) *KonversiBarangService {
	return &KonversiBarangService{repo: repo}
}

func (s *KonversiBarangService) Create(req CreateKonversiRequest) (*models.KonversiStok, error) {

	// Generate nomor transaksi
	nomorTransaksi := utils.GenerateID(config.DB, "KON", true)

	// Prepare transaksi header
	transaksi := &models.KonversiStok{
		NoKonversi: nomorTransaksi,
		Tanggal:    req.Tanggal,
	}

	// Prepare detail transaksi
	details := make([]models.KonversiStokDetail, len(req.Details))

	for i, d := range req.Details {
		details[i] = models.KonversiStokDetail{
			KodeBarang: d.KodeBarang,
			NoKonversi: d.NoKonversi,
			Quantity:   d.Quantity,
			Arah:       d.Arah,
		}
	}

	// Execute transaction
	err := s.repo.CreateTransaksiAndUpdateStock(transaksi, details)
	if err != nil {
		return nil, err
	}

	return transaksi, nil
}

func (s *KonversiBarangService) GetByNomor(nomorTransaksi string) (*models.KonversiStok, []models.KonversiStokDetail, error) {
	return s.repo.FindByNomor(nomorTransaksi)
}
