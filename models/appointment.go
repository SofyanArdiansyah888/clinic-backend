package models

import (
	"time"
)

type Appointment struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	PasienID   uint      `json:"pasien_id"`
	DokterID   uint      `json:"dokter_id"`
	CabangID   uint      `json:"cabang_id"`
	Tanggal    time.Time `json:"tanggal"`
	JamMulai   string    `json:"jam_mulai"`
	JamSelesai string    `json:"jam_selesai"`
	Status     string    `json:"status" gorm:"type:varchar(20);default:'pending'"` // pending, confirmed, cancelled, completed
	Keterangan string    `json:"keterangan" gorm:"type:text"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Pasien     Pasien    `gorm:"foreignKey:PasienID" json:"pasien"`
	Staff      Staff     `gorm:"foreignKey:DokterID" json:"dokter"`
	Cabang     Cabang    `gorm:"foreignKey:CabangID" json:"cabang"`
}
