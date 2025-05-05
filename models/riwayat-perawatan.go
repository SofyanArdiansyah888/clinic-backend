package models

import (
	customtypes "backend/models/customTypes"
	"time"
)

type RiwayatPerawatan struct {
	ID          int                  `gorm:"primary_key" json:"id"`
	IDAntrian   int64                `gorm:"not null" json:"id_antrian"`
	IDPasien    int64                `gorm:"not null" json:"id_pasien"`
	IDStaff     int64                `gorm:"not null" json:"id_staff"`
	NoPerawatan string               `gorm:"unique not null" json:"no-perawatan"`
	Pasien      Pasien               `gorm:"foreignKey:IDPasien" json:"pasien"`
	Staff       Staff                `gorm:"foreignKey:IDStaff" json:"staff"`
	Antrian     Antrian              `gorm:"foreignKey:IDAntrian" json:"antrian"`
	Tanggal     customtypes.DateTime `gorm:"not null" json:"tanggal"`
	CreatedAt   time.Time            `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time            `gorm:"autoUpdateTime" json:"updated_at"`
}
