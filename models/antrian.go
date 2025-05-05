package models

import (
	customtypes "backend/models/customTypes"
	"time"
)

type Status string

const (
	AntrianPasien    Status = "pasien"
	AntrianPerawatan Status = "perawatan"
	AntrianSelesai   Status = "selesai"
)

type Antrian struct {
	ID        int64                `gorm:"primaryKey;" json:"id"`
	IDPasien  int64                `gorm:"not null" json:"id_pasien"`
	IDStaff   int64                `gorm:"not null" json:"id_staff"`
	Pasien    Pasien               `gorm:"foreignKey:IDPasien" json:"pasien"`
	Staff     Staff                `gorm:"foreignKey:IDStaff" json:"staff"`
	Tanggal   customtypes.DateTime `gorm:"not null" json:"tanggal"`
	NoAntrian string               `json:"no_antrian"`
	Barcode   string               `json:"barcode"`
	Name      string               `json:"name"`
	Foto      string               `gorm:"null" json:"foto"`
	Status    Status               `json:"status" default:"pasien"`
	CreatedAt time.Time            `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time            `gorm:"autoUpdateTime" json:"updated_at"`
}
