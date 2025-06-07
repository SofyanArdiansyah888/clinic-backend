package penjualanBarang

import (
	customtypes "backend/models/customTypes"
)

type CreatePenjualanRequest struct {
	TanggalTransaksi customtypes.DateOnly           `json:"tanggal_transaksi" validate:"required"`
	TanggalJT        customtypes.DateOnly           `json:"tanggal_jt"`
	NoReferensi      string                         `json:"no_referensi"`
	NoCustomer       string                         `json:"no_customer"`
	JenisPembayaran  string                         `json:"jenis_pembayaran"`
	MetodePembayaran string                         `json:"metode_pembayaran"`
	IDCabang         uint                           `json:"id_cabang"`
	Details          []CreatePenjualanDetailRequest `json:"details" validate:"required,min=1"`
}
