package penjualanBarang

import (
	"backend/config"
	"backend/models"
	"backend/utils"
	"errors"
)

type PenjualanBarangService struct {
	repo *PenjualanBarangRepository
}

func NewPenjualanBarangService(repo *PenjualanBarangRepository) *PenjualanBarangService {
	return &PenjualanBarangService{repo: repo}
}

func (s *PenjualanBarangService) Create(req CreatePenjualanRequest) (*models.TransaksiBarang, error) {
	idCabang := 1
	// Validate barang existence and prepare stock updates
	stockUpdates := make(map[string]int)
	for _, detail := range req.Details {
		// Check if barang exists
		stok, err := s.repo.GetStokByKodeBarang(detail.KodeBarang)
		if err != nil {
			return nil, errors.New("barang not found: " + detail.KodeBarang)
		}

		// Calculate new stock
		newStok := stok - detail.Jumlah // Changed from + to - for sales
		if newStok < 0 {
			return nil, errors.New("insufficient stock for: " + detail.KodeBarang)
		}
		stockUpdates[detail.KodeBarang] = newStok
	}

	// Generate nomor transaksi
	nomorTransaksi := utils.GenerateID(config.DB, "INV", true)

	// Prepare transaksi header
	transaksi := &models.TransaksiBarang{
		NomorTransaksi:   nomorTransaksi,
		TanggalTransaksi: req.TanggalTransaksi,
		TanggalJT:        req.TanggalJT,
		NomorReferensi1:  req.NomorReferensi1,
		NomorReferensi2:  req.NomorReferensi2,
		NomorReferensi3:  req.NomorReferensi3,
		Tipe:             "keluar", // Changed from masuk to keluar
		JenisPembayaran:  req.JenisPembayaran,
		MetodePembayaran: req.MetodePembayaran,
		IDCabang:         uint(idCabang),
	}

	// Prepare detail transaksi
	details := make([]models.TransaksiBarangDetail, len(req.Details))
	var totalHarga float64

	for i, d := range req.Details {
		subtotal := d.Harga * float64(d.Jumlah)
		subtotal = subtotal - d.Diskon + d.PPN + d.Ongkir
		totalHarga += subtotal

		details[i] = models.TransaksiBarangDetail{
			KodeBarang:     d.KodeBarang,
			Harga:          d.Harga,
			Jumlah:         d.Jumlah,
			Diskon:         d.Diskon,
			PPN:            d.PPN,
			Ongkir:         d.Ongkir,
			Tipe:           "keluar", // Changed from masuk to keluar
			IDCabang:       uint(idCabang),
			NomorTransaksi: nomorTransaksi,
		}
	}

	transaksi.TotalHarga = totalHarga

	// Execute transaction
	err := s.repo.CreateTransaksiAndUpdateStock(transaksi, details, stockUpdates)
	if err != nil {
		return nil, err
	}

	return transaksi, nil
}

func (s *PenjualanBarangService) GetByNomor(nomorTransaksi string) (*models.TransaksiBarang, []models.TransaksiBarangDetail, error) {
	return s.repo.FindByNomor(nomorTransaksi)
}
