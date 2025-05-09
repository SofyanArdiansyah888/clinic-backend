package models

import (
	"time"
)

type Perusahaan struct {
	ID             int       `gorm:"primary_key" json:"id"`
	NamaPerusahaan string    `gorm:"string" json:"nama_perusahaan"`
	Alamat         string    `gorm:"type:text;null" json:"alamat,omitempty"`
	Telepon        string    `gorm:"type:string;null" json:"telepon,omitempty"`
	SetupCetakan   string    `gorm:"text" json:"setup_cetakan"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
