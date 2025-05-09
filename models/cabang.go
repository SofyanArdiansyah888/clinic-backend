package models

import (
	"time"
)

type Cabang struct {
	ID           int        `gorm:"primary_key" json:"id"`
	IDPerusahaan int64      `gorm:"not null" json:"id_perusahaan"`
	NamaCabang   string     `gorm:"string" json:"nama_cabang"`
	NoCabang     string     `gorm:"unique" json:"no_cabang" validate:"required"`
	Alamat       string     `gorm:"type:text;null" json:"alamat,omitempty"`
	Telepon      string     `gorm:"type:string;null" json:"telepon,omitempty"`
	Perusahaan   Perusahaan `gorm:"foreignKey:IDPerusahaan" json:"perusahaan"`
	CreatedAt    time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}
