package models

import "time"

type AppointmentTreatment struct {
	ID            uint `gorm:"primaryKey" json:"id"`
	AppointmentID uint `json:"appointment_id"`
	TreatmentID   uint `json:"treatment_id"`
	StaffID       uint `json:"staff_id"`
	Quantity      int  `json:"quantity"`
	// Appointment   Appointment `gorm:"foreignKey:AppointmentID" json:"appointment"`
	// Treatment     Treatment   `gorm:"foreignKey:TreatmentID" json:"treatment"`
	// Staff         Staff       `gorm:"foreignKey:StaffID" json:"staff"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
