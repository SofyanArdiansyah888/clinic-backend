package appointment

type CreateAppointmentDetailRequest struct {
	ID            uint `json:"id"`
	AppointmentID uint `json:"appointment_id" validate:"required"`
	TreatmentID   uint `json:"treatment_id" validate:"required"`
	StaffID       uint `json:"staff_id" validate:"required"`
	Quantity      int  `json:"quantity" validate:"required,gt=0"`
}
