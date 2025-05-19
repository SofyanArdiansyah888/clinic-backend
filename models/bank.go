package models

import (
	"time"
)

type JenisBank string

const (
	BANK   JenisBank = "bank"
	EMONEY JenisBank = "e-money"
)

type Bank struct {
	ID         int       `gorm:"primary_key" json:"id"`
	IDCabang   uint      `gorm:"not null" json:"id_cabang"`
	Cabang     Cabang    `gorm:"foreignKey:IDCabang" json:"cabang"`
	NoBank     string    `gorm:"unique" json:"no_bank" validate:"required"`
	NamaBank   string    `gorm:"string" json:"nama_bank"`
	JenisBank  JenisBank `json:"jenis_bank" validate:"required"`
	SaldoAwal  float64   `gorm:"type:float;default:0" json:"saldo_awal"`
	NoRekening string    `gorm:"string" json:"no_rekening"`
	AtasNama   string    `gorm:"string" json:"atas_nama"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
}
