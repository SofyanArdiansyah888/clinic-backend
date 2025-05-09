package models

import (
	"time"
)

type Bank struct {
	ID        int       `gorm:"primary_key" json:"id"`
	IDCabang  int64     `gorm:"not null" json:"id_cabang"`
	NamaBank  string    `gorm:"string" json:"nama_bank"`
	NoBank    string    `gorm:"unique" json:"no-bank" validate:"required"`
	SaldoAwal float64   `gorm:"type:float;default:0" json:"saldo_awal"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
