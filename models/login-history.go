package models

import (
	"time"
)

type LoginHistory struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"primaryKey" json:"user_id"`
	Username  string    `json:"username"`
	Waktu     int64     `gorm:"datetime" json:"waktu"`
	Asal      string    `json:"asal"`
	Notes     string    `gorm:"text" json:"notes"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
