package models

import (
	"time"

	"gorm.io/gorm"
)

type StokMovement struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey" json:"id"`
	KodeBarang string `gorm:"type:varchar(100);not null;index" json:"kode_barang"` // ← ensure it's indexed
	// Barang     Barang `gorm:"foreignKey:KodeBarang;references:KodeBarang" json:"barang"`
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
