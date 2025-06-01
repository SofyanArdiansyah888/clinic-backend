package pembelianBarang

import (
	"backend/config"
	"backend/models"
	"backend/utils"
	"errors"
)

type PembelianBarangService struct {
	repo *PembelianBarangRepository
}

func NewPembelianBarangService(repo *PembelianBarangRepository) *PembelianBarangService {
	return &PembelianBarangService{repo: repo}
}

func (s *PembelianBarangService) Create(req CreatePembelianRequest) (*models.TransaksiBarang, error) {
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
		newStok := stok + detail.Jumlah
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
		Tipe:             "masuk",
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
			Tipe:           "masuk",
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

func (s *PembelianBarangService) GetByNomor(nomorTransaksi string) (*models.TransaksiBarang, []models.TransaksiBarangDetail, error) {
	return s.repo.FindByNomor(nomorTransaksi)
}
