package pembelianBarang

import (
	"backend/config"
	"backend/models"
	"backend/utils"
)

type PembelianBarangService struct {
	repo *PembelianBarangRepository
}

func NewPembelianBarangService(repo *PembelianBarangRepository) *PembelianBarangService {
	return &PembelianBarangService{repo: repo}
}

func (s *PembelianBarangService) Create(req CreatePembelianRequest) (*models.Pembelian, error) {
	idCabang := 1

	// Generate nomor transaksi
	nomorTransaksi := utils.GenerateID(config.DB, "INV", true)

	// Prepare transaksi header
	transaksi := &models.Pembelian{
		NoTransaksi:      nomorTransaksi,
		TanggalTransaksi: req.TanggalTransaksi,
		TanggalJT:        req.TanggalJT,
		NoReferensi:      req.NoReferensi,
		Tipe:             "masuk",
		JenisPembayaran:  req.JenisPembayaran,
		MetodePembayaran: req.MetodePembayaran,
		IDCabang:         uint(idCabang),
	}

	// Prepare detail transaksi
	details := make([]models.PembelianDetail, len(req.Details))
	var totalHarga float64

	for i, d := range req.Details {
		subtotal := d.Harga * float64(d.Jumlah)
		subtotal = subtotal - d.Diskon + d.PPN + d.Ongkir
		totalHarga += subtotal

		details[i] = models.PembelianDetail{
			KodeBarang:  d.KodeBarang,
			Harga:       d.Harga,
			Jumlah:      d.Jumlah,
			Diskon:      d.Diskon,
			PPN:         d.PPN,
			Ongkir:      d.Ongkir,
			Tipe:        "masuk",
			IDCabang:    uint(idCabang),
			NoTransaksi: nomorTransaksi,
		}
	}

	transaksi.TotalHarga = totalHarga

	// Execute transaction
	err := s.repo.CreateTransaksiAndUpdateStock(transaksi, details)
	if err != nil {
		return nil, err
	}

	return transaksi, nil
}

func (s *PembelianBarangService) GetByNomor(nomorTransaksi string) (*models.Pembelian, []models.PembelianDetail, error) {
	return s.repo.FindByNomor(nomorTransaksi)
}
