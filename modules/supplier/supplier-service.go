package supplier

import (
	"backend/models"
)

type ISupplierService interface {
	GetAll() ([]models.Supplier, error)
	GetByID(id uint) (*models.Supplier, error)
	Create(data *models.Supplier) error
	Update(id uint, data *models.Supplier) error
	Delete(id uint) error
}

type supplierService struct {
	repo ISupplierRepository
}

func NewSupplierService(repo ISupplierRepository) *supplierService {
	return &supplierService{repo: repo}
}

func (s *supplierService) GetAll() ([]models.Supplier, error) {
	return s.repo.FindAll()
}

func (s *supplierService) GetByID(id uint) (*models.Supplier, error) {
	return s.repo.FindByID(id)
}

func (s *supplierService) Create(data *models.Supplier) error {
	//data.NoRM = utils.GenerateID(config.DB, "RMD", true)
	//data.NoMember = utils.GenerateID(config.DB, "MEM", true)
	return s.repo.Create(data)
}

func (s *supplierService) Update(id uint, data *models.Supplier) error {
	staff, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	data.NoSupplier = staff.NoSupplier
	data.ID = staff.ID
	return s.repo.Update(data)
}

func (s *supplierService) Delete(id uint) error {
	return s.repo.Delete(id)
}
