package produksiBarang

type CreateProduksiDetailRequest struct {
	KodeBarang string  `json:"kode_barang" validate:"required"`
	IDProduksi string  `json:"id_produksi" validate:"required"`
	Quantity   float64 `json:"quantity" validate:"required,gt=0"`
	Arah       string  `json:"arah" validate:"required"`
}
