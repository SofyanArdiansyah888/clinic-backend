package models

import (
	"time"
)

type TransaksiBarang struct {
	ID                uint      `json:"id" gorm:"primaryKey"`
	NomorTransaksi    string    `json:"nomor_transaksi" gorm:"type:varchar(100);not null"`
	TanggalTransaksi  time.Time `json:"tanggal_transaksi" gorm:"not null"`
	TanggalJT        time.Time `json:"tanggal_jt" gorm:"default:null"`
	NomorReferensi1   string    `json:"nomor_referensi_1" gorm:"type:varchar(100)"`
	NomorReferensi2   string    `json:"nomor_referensi_2" gorm:"type:varchar(100)"`
	NomorReferensi3   string    `json:"nomor_referensi_3" gorm:"type:varchar(100)"`
	Tipe             string    `json:"tipe" gorm:"type:varchar(20);not null"` // keluar, masuk
	JenisPembayaran  string    `json:"jenis_pembayaran" gorm:"type:varchar(50);not null"`
	MetodePembayaran string    `json:"metode_pembayaran" gorm:"type:varchar(50);not null"`
	TotalDiskon       float64   `json:"total_diskon" gorm:"type:decimal(15,2);not null;default:0"`
	TotalPPN       float64   `json:"total_ppn" gorm:"type:decimal(15,2);not null;default:0"`
	TotalOngkir    float64   `json:"total_ongkir" gorm:"type:decimal(15,2);not null;default:0"`
	TotalHarga       float64   `json:"total_harga" gorm:"type:decimal(15,2);not null;default:0"`
	TotalHPP       float64   `json:"total_hpp" gorm:"type:decimal(15,2);not null;default:0"`
	IDCabang         uint      `json:"id_cabang" gorm:"not null"`
	Cabang           Cabang    `json:"cabang" gorm:"foreignKey:IDCabang"`
}