package models

import (
	"time"
)

type Barang struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	SupplierID uint      `gorm:"primaryKey" json:"supplier_id"`
	NamaBarang string    `gorm:"type:text;not null" json:"nama_barang"`
	KodeBarang string    `gorm:"type:varchar(100);not null" json:"kode_barang"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
