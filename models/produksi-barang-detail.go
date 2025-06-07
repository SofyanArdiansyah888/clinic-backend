package models

import "time"

type ProduksiBarangDetail struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	NoProduksi string    `json:"no_produksi"`
	KodeBarang string    `json:"kode_barang"`
	Quantity   float64   `json:"quantity"`
	Arah       string    `json:"arah"` // keluar atau masuk
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
