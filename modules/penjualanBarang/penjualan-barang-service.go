package penjualanBarang

import (
	"backend/config"
	"backend/models"
	"backend/utils"
)

type PenjualanBarangService struct {
	repo *PenjualanBarangRepository
}

func NewPenjualanBarangService(repo *PenjualanBarangRepository) *PenjualanBarangService {
	return &PenjualanBarangService{repo: repo}
}

func (s *PenjualanBarangService) Create(req CreatePenjualanRequest) (*models.Penjualan, error) {
	idCabang := 1

	// Generate nomor transaksi
	nomorTransaksi := utils.GenerateID(config.DB, "INV", true)

	// Prepare transaksi header
	transaksi := &models.Penjualan{
		NoTransaksi:      nomorTransaksi,
		TanggalTransaksi: req.TanggalTransaksi,
		TanggalJT:        req.TanggalJT,
		NoReferensi:      req.NoReferensi,
		NoCustomer:       req.NoCustomer,
		Tipe:             "keluar", // Changed from masuk to keluar
		JenisPembayaran:  req.JenisPembayaran,
		MetodePembayaran: req.MetodePembayaran,
		IDCabang:         uint(idCabang),
	}

	// Prepare detail transaksi
	details := make([]models.PenjualanDetail, len(req.Details))
	var totalHarga float64

	for i, d := range req.Details {
		subtotal := d.Harga * float64(d.Jumlah)
		subtotal = subtotal - d.Diskon + d.PPN + d.Ongkir
		totalHarga += subtotal

		details[i] = models.PenjualanDetail{
			KodeBarang:  d.KodeBarang,
			Harga:       d.Harga,
			Jumlah:      d.Jumlah,
			Diskon:      d.Diskon,
			PPN:         d.PPN,
			Ongkir:      d.Ongkir,
			Tipe:        "keluar", // Changed from masuk to keluar
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

func (s *PenjualanBarangService) GetByNomor(nomorTransaksi string) (*models.Penjualan, []models.PenjualanDetail, error) {
	return s.repo.FindByNomor(nomorTransaksi)
}
