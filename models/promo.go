package models

import (
	"time"
)

type Promo struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	NamaPromo string    `gorm:"not null" json:"nama_promo"`
	KodePromo string    `gorm:"unique" json:"kode_promo"`
	Deskripsi string    `json:"deskripsi"`
	Diskon    float64   `json:"diskon"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
