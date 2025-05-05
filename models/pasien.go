package models

import (
	"backend/models/customTypes"
	"time"
)

type Pasien struct {
	ID               int                  `gorm:"primary_key" json:"id"`
	TanggalDaftar    customTypes.DateOnly `gorm:"type:date" json:"tanggal_daftar" validate:"required"`
	NamaPasien       string               `json:"nama_pasien" validate:"required"`
	NoRM             string               `gorm:"unique" json:"no_rm" validate:"required"`
	NoMember         string               `gorm:"unique" json:"no_member" validate:"required"`
	NoIdentitas      string               `gorm:"unique" json:"no_identitas" validate:"required"`
	Agama            string               `json:"agama"`
	Alamat           string               `gorm:"text" json:"alamat"`
	Kota             string               `json:"kota"`
	JenisKelamin     string               `json:"jenis_kelamin" validate:"required,oneof=laki-laki perempuan"`
	Pekerjaan        string               `json:"pekerjaan"`
	StatusPernikahan string               `json:"status_pernikahan"`
	TempatLahir      string               `json:"tempat_lahir"`
	TanggalLahir     customTypes.DateOnly `gorm:"type:date" json:"tanggal_lahir" validate:"required"`
	Email            string               `json:"email" validate:"omitempty,email"`
	NomorHP          string               `json:"nomor_hp" validate:"required"`
	Catatan          string               `gorm:"text" json:"catatan"`
	CreatedAt        time.Time            `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time            `gorm:"autoUpdateTime" json:"updated_at"`
}
