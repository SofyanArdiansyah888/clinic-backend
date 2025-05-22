package models

import (
	"time"
)

type LoginHistory struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	IDUser    uint      `gorm:"primaryKey" json:"id_user"`
	Username  string    `json:"username"`
	Waktu     int64     `gorm:"datetime" json:"waktu"`
	Asal      string    `json:"asal"`
	Notes     string    `gorm:"text" json:"notes"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
