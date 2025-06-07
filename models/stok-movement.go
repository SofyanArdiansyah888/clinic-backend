package models

import (
	"time"
)

type StokMovement struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	KodeBarang string `gorm:"not null" json:"kode_barang"`
	// Barang     Barang `gorm:"foreignKey:KodeBarang" json:"barang"`
	Quantity int    `gorm:"not null" json:"quantity"`
	Jenis    string `gorm:"not null" json:"jenis"`
	//  pembelian, penjualan, penyesuaian, konversi masuk, konversi keluar,
	//  produksi masuk, produksi keluar, transfer masuk, transfer keluar,
	//  return pembelian, retur penjualan
	KodeReferensi string    `gorm:"not null" json:"kode_referensi"`
	Keterangan    string    `json:"keterangan"`
	CreatedAt     time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt     time.Time `gorm:"not null" json:"updated_at"`
}
