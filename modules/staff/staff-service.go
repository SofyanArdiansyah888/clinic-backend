package staff

import (
	"backend/config"
	"backend/models"
	"backend/utils"
)

type IStaffService interface {
	GetAll() ([]models.Staff, error)
	GetByID(id uint) (*models.Staff, error)
	Create(data *models.Staff) error
	Update(id uint, data *models.Staff) error
	Delete(id uint) error
}

type StaffService struct {
	repo IStaffRepository
}

func NewStaffService(repo IStaffRepository) *StaffService {
	return &StaffService{repo: repo}
}

func (s *StaffService) GetAll() ([]models.Staff, error) {
	return s.repo.FindAll()
}

func (s *StaffService) GetByID(id uint) (*models.Staff, error) {
	return s.repo.FindByID(id)
}

func (s *StaffService) Create(data *models.Staff) error {
	data.NoStaff = utils.GenerateID(config.DB, "STF", true)
	return s.repo.Create(data)
}

func (s *StaffService) Update(id uint, data *models.Staff) error {
	staff, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	data.NoStaff = staff.NoStaff
	data.ID = staff.ID
	return s.repo.Update(data)
}

func (s *StaffService) Delete(id uint) error {
	return s.repo.Delete(id)
}
