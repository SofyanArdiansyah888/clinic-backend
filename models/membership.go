package models

import (
	"time"
)

type Membership struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	PasienID        uint      `json:"pasien_id"`
	NomorMembership string    `gorm:"type:varchar(50);unique" json:"nomor_membership"`
	TipeMembership  string    `gorm:"type:varchar(50)" json:"tipe_membership"` // silver, gold, platinum
	TanggalMulai    time.Time `json:"tanggal_mulai"`
	TanggalBerakhir time.Time `json:"tanggal_berakhir"`
	Status          string    `gorm:"type:varchar(20);default:'active'" json:"status"` // active, expired, cancelled
	Poin            int       `gorm:"default:0" json:"poin"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Pasien          Pasien    `gorm:"foreignKey:PasienID" json:"pasien"`
}
