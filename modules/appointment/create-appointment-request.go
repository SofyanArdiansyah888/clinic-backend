package appointment

type CreateAppointmentRequest struct {
	PasienID   uint                             `json:"pasien_id" validate:"required"`
	StaffID    uint                             `json:"staff_id" validate:"required"`
	CabangID   uint                             `json:"cabang_id" validate:"required"`
	Tanggal    string                           `json:"tanggal" validate:"required,datetime=2006-01-02"`
	JamMulai   string                           `json:"jam_mulai" validate:"required"`
	JamSelesai string                           `json:"jam_selesai" validate:"required"`
	Status     string                           `json:"status" validate:"required"`
	Details    []CreateAppointmentDetailRequest `json:"details" validate:"required,dive"`
}
