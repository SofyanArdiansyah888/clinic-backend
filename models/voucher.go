package models

import (
	"time"

	"gorm.io/gorm"
)

type Voucher struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	KodeVoucher     string         `json:"kode_voucher" gorm:"uniqueIndex;not null"`
	NamaVoucher     string         `json:"nama_voucher" gorm:"not null"`
	Deskripsi       string         `json:"deskripsi"`
	TipeDiskon      string         `json:"tipe_diskon" gorm:"not null"` // percentage or fixed
	NilaiDiskon     float64        `json:"nilai_diskon" gorm:"not null"`
	MinimumOrder    float64        `json:"minimum_order"`
	MaksimumDiskon  float64        `json:"maksimum_diskon"`
	Kuota           int            `json:"kuota"`
	TanggalMulai    time.Time      `json:"tanggal_mulai" gorm:"not null"`
	TanggalBerakhir time.Time      `json:"tanggal_berakhir" gorm:"not null"`
	Status          string         `json:"status" gorm:"default:'active'"` // active, inactive, expired
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
