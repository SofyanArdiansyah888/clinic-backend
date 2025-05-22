package models

import (
	"time"
)

type TemplateConcern struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Nama      string    `gorm:"not null" json:"nama"`
	Jenis     string    `json:"jenis"`
	Template  string    `json:"template"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}