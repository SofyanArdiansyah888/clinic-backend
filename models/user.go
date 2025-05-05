package models

import (
	"time"
)

type Role string

const (
	Administrator Role = "administrator"
	Management    Role = "management"
	Gudang        Role = "gudang"
	Marketing     Role = "marketing"
	Resepsionis   Role = "resepsionis"
	Klinik        Role = "klinik"
	Apotik        Role = "apotik"
	Kasir         Role = "kasir"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CabangID  uint      `gorm:"primaryKey" json:"cabang_id"`
	NoUser    string    `gorm:"unique" json:"nomor_user"`
	Username  string    `gorm:"unique" json:"username"`
	Nama      string    `json:"nama"`
	Password  string    `json:"password"`
	Akses     string    `json:"akses"`
	Role      Role      `json:"role"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
