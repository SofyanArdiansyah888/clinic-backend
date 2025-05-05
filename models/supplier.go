package models

import (
	"time"
)

type Supplier struct {
	ID         int       `gorm:"primary_key" json:"id"`
	Nama       string    `json:"nama" validate:"required"`
	NoSupplier string    `gorm:"unique" json:"no_supplier" validate:"required"`
	Telepon    string    `json:"telepon" validate:"required"`
	Alamat     string    `gorm:"text" json:"alamat"`
	Kota       string    `json:"kota"`
	Pic        string    `json:"pic"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
