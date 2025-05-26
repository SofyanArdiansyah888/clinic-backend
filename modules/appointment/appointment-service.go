package appointment

import (
	"backend/models"
	"errors"
	"time"
)

type IAppointmentService interface {
	GetAll() ([]models.Appointment, error)
	GetByID(id uint) (*models.Appointment, error)
	Create(data *models.Appointment) error
	Update(id uint, data *models.Appointment) error
	Delete(id uint) error
}

type appointmentService struct {
	repo IAppointmentRepository
}

func NewAppointmentService(repo IAppointmentRepository) *appointmentService {
	return &appointmentService{repo: repo}
}

func (s *appointmentService) GetAll() ([]models.Appointment, error) {
	return s.repo.FindAll()
}

func (s *appointmentService) GetByID(id uint) (*models.Appointment, error) {
	return s.repo.FindByID(id)
}

func (s *appointmentService) Create(data *models.Appointment) error {
	// Validate required fields
	if data.PasienID == 0 || data.DokterID == 0 || data.CabangID == 0 {
		return errors.New("pasien, dokter, and cabang are required")
	}

	// Validate appointment time
	if data.Tanggal.Before(time.Now()) {
		return errors.New("appointment date must be in the future")
	}

	// Set initial status if not provided
	if data.Status == "" {
		data.Status = "pending"
	}

	return s.repo.Create(data)
}

func (s *appointmentService) Update(id uint, data *models.Appointment) error {
	// Check if appointment exists
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	// Validate status transition
	if !isValidStatusTransition(existing.Status, data.Status) {
		return errors.New("invalid status transition")
	}

	data.ID = existing.ID
	return s.repo.Update(data)
}

func (s *appointmentService) Delete(id uint) error {
	// Check if appointment exists
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}

// Helper function to validate status transitions
func isValidStatusTransition(currentStatus, newStatus string) bool {
	if newStatus == "" {
		return true // Allow empty status in updates
	}

	validTransitions := map[string][]string{
		"pending":   {"confirmed", "cancelled"},
		"confirmed": {"completed", "cancelled"},
		"cancelled": {},
		"completed": {},
	}

	allowedStatuses, exists := validTransitions[currentStatus]
	if !exists {
		return false
	}

	for _, status := range allowedStatuses {
		if status == newStatus {
			return true
		}
	}

	return false
}
