package konversiBarang

type CreateKonversiDetailRequest struct {
	KodeBarang string  `json:"kode_barang" validate:"required"`
	IDKonversi string  `json:"id_konversi" validate:"required"`
	Quantity   float64 `json:"quantity" validate:"required,gt=0"`
	Arah       string  `json:"arah" validate:"required"`
}
