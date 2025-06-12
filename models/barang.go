package models

import (
	"time"

	"gorm.io/gorm"
)

type Barang struct {
	gorm.Model
	ID           uint   `gorm:"primaryKey" json:"id"`
	NamaBarang   string `gorm:"type:text;not null" json:"nama_barang"`
	KodeBarang   string `gorm:"type:varchar(100);not null;unique" json:"kode_barang"`
	Satuan       string `gorm:"type:varchar(100);not null" json:"satuan"`
	JenisBarang  string `gorm:"type:varchar(100);not null" json:"jenis_barang"`
	LokasiBarang string `gorm:"type:varchar(100);not null;default:gudang" json:"lokasi_barang"` // gudang, apotek
	MinStock     int    `gorm:"type:int;not null" json:"min_stock"`
	// Stock         *int           `gorm:"type:int" json:"stock"`
	// StokMovements []StokMovement `gorm:"foreignKey:KodeBarang;references:KodeBarang" json:"stok_movements"`
	HargaBeli int       `gorm:"type:int;not null;default:0" json:"harga_beli"`
	HargaJual int       `gorm:"type:int;not null;default:0" json:"harga_jual"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
