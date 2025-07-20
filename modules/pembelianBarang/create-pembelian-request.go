package pembelianBarang

import (
	customtypes "backend/models/customTypes"
)

type CreatePembelianRequest struct {
	TanggalTransaksi customtypes.DateOnly           `json:"tanggal_transaksi" validate:"required"`
	TanggalJT        customtypes.DateOnly           `json:"tanggal_jt"`
	IDReferensi      string                         `json:"id_referensi"`
	IDSupplier       string                         `json:"id_supplier"`
	JenisPembayaran  string                         `json:"jenis_pembayaran" validate:"required"`
	MetodePembayaran string                         `json:"metode_pembayaran" validate:"required"`
	IDCabang         uint                           `json:"id_cabang"`
	Details          []CreatePembelianDetailRequest `json:"details" validate:"required,min=1"`
}
