package penjualanBarang

type CreatePenjualanDetailRequest struct {
	KodeBarang string  `json:"kode_barang" validate:"required"`
	Harga      float64 `json:"harga" validate:"required,gt=0"`
	Jumlah     int     `json:"jumlah" validate:"required,gt=0"`
	Diskon     float64 `json:"diskon" validate:"gte=0"`
	PPN        float64 `json:"ppn" validate:"gte=0"`
	Ongkir     float64 `json:"ongkir" validate:"gte=0"`
}
