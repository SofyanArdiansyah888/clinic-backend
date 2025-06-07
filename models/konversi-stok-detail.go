package models

import "time"

type KonversiStokDetail struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	NoKonversi string    `json:"no_konversi"`
	KodeBarang string    `json:"kode_barang"`
	Quantity   float64   `json:"quantity"`
	Arah       string    `json:"arah"` // keluar atau masuk
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
