package models

import (
	customTypes "backend/models/customTypes"
	"time"
)

type StokOpname struct {
	ID           uint                 `json:"id" gorm:"primaryKey"`
	Tanggal      customTypes.DateOnly `json:"date"`
	KodeBarang   string               `json:"kode_barang" gorm:"type:varchar(100);not null"`
	NoStokOpname string               `gorm:"type:varchar(100);not null;unique" json:"no_stok_opname"`
	Barang       Barang               `json:"barang" gorm:"foreignKey:kode_barang"`
	StokSistem   int                  `json:"stok_sistem"`
	StokRiil     int                  `json:"stok_riil"`
	Alasan       string               `json:"alasan"`
	CreatedAt    time.Time            `json:"created_at"`
	UpdatedAt    time.Time            `json:"updated_at"`
}
