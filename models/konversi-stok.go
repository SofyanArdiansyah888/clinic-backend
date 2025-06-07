package models

import (
	customTypes "backend/models/customTypes"
	"time"
)

type KonversiStok struct {
	ID         uint                 `json:"id" gorm:"primaryKey"`
	NoKonversi string               `json:"no_konversi" gorm:"not null;unique"`
	Tanggal    customTypes.DateOnly `json:"tanggal" gorm:"not null"`
	CreatedAt  time.Time            `json:"created_at"`
	UpdatedAt  time.Time            `json:"updated_at"`
}
