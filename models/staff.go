package models

import (
	"time"
)

type Jabatan string

const (
	Dokter     Jabatan = "dokter"
	Beautician Jabatan = "beautician"
	Nurse      Jabatan = "nurse"
	Sales      Jabatan = "sales"
	Karyawan   Jabatan = "karyawan"
)

type Level string

const (
	Junior    Level = "1"
	Senior    Level = "2"
	Spesialis Level = "3"
)

type Staff struct {
	ID        int       `gorm:"primary_key" json:"id"`
	NoStaff   string    `gorm:"unique" json:"no_staff"`
	Jabatan   Jabatan   `json:"jabatan" validate:"required"`
	Nama      string    `json:"nama" validate:"required"`
	Telepon   string    `json:"telepon"`
	Level     Level     `json:"level" validate:"required"`
	Alamat    string    `gorm:"text" json:"alamat"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
