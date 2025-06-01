package penjualanBarang

import (
	customtypes "backend/models/customTypes"
)

type CreatePenjualanRequest struct {
	TanggalTransaksi customtypes.DateOnly           `json:"tanggal_transaksi" validate:"required"`
	TanggalJT        customtypes.DateOnly           `json:"tanggal_jt"`
	NomorReferensi1  string                         `json:"nomor_referensi_1"`
	NomorReferensi2  string                         `json:"nomor_referensi_2"`
	NomorReferensi3  string                         `json:"nomor_referensi_3"`
	JenisPembayaran  string                         `json:"jenis_pembayaran"`
	MetodePembayaran string                         `json:"metode_pembayaran"`
	IDCabang         uint                           `json:"id_cabang"`
	Details          []CreatePenjualanDetailRequest `json:"details" validate:"required,min=1"`
}
