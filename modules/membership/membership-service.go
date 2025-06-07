package membership

import (
	"backend/config"
	"backend/models"
	"backend/utils"
	"errors"
	"time"
)

type IMembershipService interface {
	GetAll() ([]models.Membership, error)
	GetByID(id uint) (*models.Membership, error)
	GetByPasienID(pasienID uint) ([]models.Membership, error)
	Create(data *models.Membership) error
	Update(id uint, data *models.Membership) error
	Delete(id uint) error
	UpdatePoints(id uint, points int) error
}

type membershipService struct {
	repo IMembershipRepository
}

func NewMembershipService(repo IMembershipRepository) *membershipService {
	return &membershipService{repo: repo}
}

func (s *membershipService) GetAll() ([]models.Membership, error) {
	return s.repo.FindAll()
}

func (s *membershipService) GetByID(id uint) (*models.Membership, error) {
	return s.repo.FindByID(id)
}

func (s *membershipService) GetByPasienID(pasienID uint) ([]models.Membership, error) {
	return s.repo.FindByPasienID(pasienID)
}

func (s *membershipService) Create(data *models.Membership) error {
	// Validate required fields
	if data.PasienID == 0 || data.TipeMembership == "" {
		return errors.New("pasien and membership type are required")
	}

	// Validate membership type
	if !isValidMembershipType(data.TipeMembership) {
		return errors.New("invalid membership type")
	}

	// Generate membership number
	data.NoMembership = utils.GenerateID(config.DB, "MBR", true)

	// Set default status if not provided
	if data.Status == "" {
		data.Status = "active"
	}

	// Set membership duration if not provided
	if data.TanggalMulai.IsZero() {
		data.TanggalMulai = time.Now()
	}
	if data.TanggalBerakhir.IsZero() {
		data.TanggalBerakhir = data.TanggalMulai.AddDate(1, 0, 0) // 1 year validity
	}

	return s.repo.Create(data)
}

func (s *membershipService) Update(id uint, data *models.Membership) error {
	// Check if membership exists
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	// Validate membership type if changed
	if data.TipeMembership != "" && !isValidMembershipType(data.TipeMembership) {
		return errors.New("invalid membership type")
	}

	// Validate status transition if changed
	if data.Status != "" && !isValidStatusTransition(existing.Status, data.Status) {
		return errors.New("invalid status transition")
	}

	// Preserve unchangeable fields
	data.ID = existing.ID
	data.NoMembership = existing.NoMembership
	data.PasienID = existing.PasienID

	return s.repo.Update(data)
}

func (s *membershipService) Delete(id uint) error {
	// Check if membership exists
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}

func (s *membershipService) UpdatePoints(id uint, points int) error {
	membership, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	membership.Poin += points
	if membership.Poin < 0 {
		membership.Poin = 0
	}

	return s.repo.Update(membership)
}

// Helper function to validate membership type
func isValidMembershipType(membershipType string) bool {
	validTypes := map[string]bool{
		"silver":   true,
		"gold":     true,
		"platinum": true,
	}

	_, valid := validTypes[membershipType]
	return valid
}

// Helper function to validate status transitions
func isValidStatusTransition(currentStatus, newStatus string) bool {
	if newStatus == "" {
		return true // Allow empty status in updates
	}

	validTransitions := map[string][]string{
		"active":    {"expired", "cancelled"},
		"expired":   {"active"},
		"cancelled": {"active"},
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
