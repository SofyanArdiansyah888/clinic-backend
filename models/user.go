package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// Custom type for Akses
type StringArray []string

// Implement the driver.Valuer interface
func (a StringArray) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Implement the sql.Scanner interface
func (a *StringArray) Scan(value interface{}) error {
	if value == nil {
		*a = nil
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, a)
}

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
	ID       uint    `gorm:"primaryKey" json:"id"`
	IDCabang int64   `gorm:"not null" json:"id_cabang"`
	Cabang   *Cabang `gorm:"foreignKey:IDCabang;" json:"cabang"`
	NoUser   string  `gorm:"unique" json:"no_user"`
	Username string  `gorm:"unique" json:"username"`
	Nama     string  `json:"nama"`
	Password string  `json:"password"`
	//@TODO: gorm:"type:jsonb" json:"akses"`
	//Akses     StringArray `gorm:"type:jsonb" json:"akses"`
	Akses     StringArray `gorm:"type:json" json:"akses"`
	Role      Role        `json:"role"`
	CreatedAt time.Time   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
}
