package models

import (
	"time"
)

type MonthlySequence struct {
	ID        uint `gorm:"primaryKey"`
	Model     string
	YearMonth string // Format: YYMM
	Counter   int
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
