package models

import (
	"time"
)

type Cabang struct {
	ID                    uint      `gorm:"primaryKey" json:"id"`
	NoCabang              string    `gorm:"not null;unique" json:"no_cabang"`
	NamaKlinik            string    `gorm:"not null" json:"nama_klinik"`
	AlamatLengkap         string    `gorm:"not null" json:"alamat_lengkap"`
	NoTelp                string    `json:"no_telp"`
	NoHandphone           string    `json:"no_handphone"`
	EmailKlinik           string    `gorm:"unique" json:"email_klinik"`
	LatitudeLongitude     string    `json:"latitude_longitude"`
	MarginHrgJualObat     float64   `json:"margin_hrg_jual_obat"`
	TargetPemasukanHarian float64   `json:"target_pemasukan_harian"`
	NominalDiskonReseller float64   `json:"nominal_diskon_reseller"`
	CreatedAt             time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt             time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
