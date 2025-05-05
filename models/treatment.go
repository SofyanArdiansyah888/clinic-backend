package models

import (
	"time"
)

type Treatment struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	NamaTreatment     string    `gorm:"not null" json:"nama_treatment"`
	NoTreatment       string    `gorm:"not null" json:"no_treatment"`
	UnitTreatment     string    `gorm:"not null" json:"unit_treatment"`
	JenisTreatment    string    `gorm:"not null" json:"jenis_treatment"`
	WaktuPengerjaan   int       `gorm:"not null" json:"waktu_pengerjaan"` // bisa jadi time.Duration jika diperlukan
	TarifUmum         float64   `gorm:"not null" json:"tarif_umum"`
	TarifMember       float64   `gorm:"not null" json:"tarif_member"`
	FeeDokter         float64   `gorm:"not null" json:"fee_dokter"`
	FeeDokterType     string    `gorm:"not null" json:"fee_dokter_type"`
	FeeBeautician     float64   `gorm:"not null" json:"fee_beautician"`
	FeeBeauticianType string    `gorm:"not null" json:"fee_beautician_type"`
	FeeSales          float64   `gorm:"not null" json:"fee_sales"`
	FeeSalesType      string    `gorm:"not null" json:"fee_sales_type"`
	BiayaModal        float64   `gorm:"not null" json:"biaya_modal"`
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
