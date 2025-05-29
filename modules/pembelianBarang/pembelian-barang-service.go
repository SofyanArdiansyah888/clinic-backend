package pembelianBarang

import (
	"backend/models"
	"backend/utils"
	"time"
)

type PembelianBarangService struct {
	repo *PembelianBarangRepository
}

func NewPembelianBarangService(repo *PembelianBarangRepository) *PembelianBarangService {
	return &PembelianBarangService{repo: repo}
}

type CreatePembelianRequest struct {
	TanggalTransaksi  time.Time                      `json:"tanggal_transaksi" validate:"required"`
	TanggalJT        time.Time                      `json:"tanggal_jt" validate:"required"`
	NomorReferensi1   string                         `json:"nomor_referensi_1"`
	NomorReferensi2   string                         `json:"nomor_referensi_2"`
	NomorReferensi3   string                         `json:"nomor_referensi_3"`
	JenisPembayaran  string                         `json:"jenis_pembayaran" validate:"required"`
	MetodePembayaran string                         `json:"metode_pembayaran" validate:"required"`
	IDCabang         uint                           `json:"id_cabang" validate:"required"`
	Details          []CreatePembelianDetailRequest `json:"details" validate:"required,min=1"`
}

type CreatePembelianDetailRequest struct {
	KodeBarang string  `json:"kode_barang" validate:"required"`
	Harga      float64 `json:"harga" validate:"required"`
	Jumlah     int     `json:"jumlah" validate:"required"`
	Diskon     float64 `json:"diskon"`
	PPN        float64 `json:"ppn"`
	Ongkir     float64 `json:"ongkir"`
}

func (s *PembelianBarangService) Create(req CreatePembelianRequest) (*models.TransaksiBarang, []models.TransaksiBarangDetail, error) {
	// Generate nomor transaksi
	nomorTransaksi := utils.GenerateTransactionNumber("INV", req.IDCabang)

	// Hitung total
	var totalHarga float64
	details := make([]models.TransaksiBarangDetail, len(req.Details))
	for i, d := range req.Details {
		subtotal := d.Harga * float64(d.Jumlah)
		subtotal = subtotal - d.Diskon + d.PPN + d.Ongkir
		totalHarga += subtotal

		details[i] = models.TransaksiBarangDetail{
			KodeBarang: d.KodeBarang,
			Harga:      d.Harga,
			Jumlah:     d.Jumlah,
			Diskon:     d.Diskon,
			PPN:        d.PPN,
			Ongkir:     d.Ongkir,
			Tipe:       "masuk",
		}
	}

	// Buat transaksi
	transaksi := &models.TransaksiBarang{
		NomorTransaksi:   nomorTransaksi,
		TanggalTransaksi: req.TanggalTransaksi,
		TanggalJT:       req.TanggalJT,
		NomorReferensi1:  req.NomorReferensi1,
		NomorReferensi2:  req.NomorReferensi2,
		NomorReferensi3:  req.NomorReferensi3,
		Tipe:            "masuk",
		JenisPembayaran: req.JenisPembayaran,
		MetodePembayaran: req.MetodePembayaran,
		TotalHarga:      totalHarga,
		IDCabang:        req.IDCabang,
	}

	// Simpan ke database
	err := s.repo.Create(transaksi, details)
	if err != nil {
		return nil, nil, err
	}

	return transaksi, details, nil
}

func (s *PembelianBarangService) GetByNomor(nomorTransaksi string) (*models.TransaksiBarang, []models.TransaksiBarangDetail, error) {
	return s.repo.FindByNomor(nomorTransaksi)
}