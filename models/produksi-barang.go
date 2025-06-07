package models

import (
	customTypes "backend/models/customTypes"
	"time"
)

type ProduksiBarang struct {
	ID         uint                 `json:"id" gorm:"primaryKey"`
	NoProduksi uint                 `json:"no_produksi" gorm:"not null;unique"`
	Tanggal    customTypes.DateOnly `json:"tanggal" gorm:"not null"`
	CreatedAt  time.Time            `json:"created_at"`
	UpdatedAt  time.Time            `json:"updated_at"`
}
