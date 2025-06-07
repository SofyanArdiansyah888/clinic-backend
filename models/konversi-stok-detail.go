package models

import "time"

type KonversiStokDetail struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	NoKonversi uint      `json:"no_konversi"`
	KodeBarang string    `json:"kode_barang"`
	Quantity   float64   `json:"quantity"`
	Arah       string    `json:"arah"` // keluar atau masuk
	UpdatedAt  time.Time `json:"updated_at"`
}
