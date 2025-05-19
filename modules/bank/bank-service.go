package bank

import (
	"backend/config"
	"backend/models"
	"backend/utils"
)

type IBankService interface {
	GetAll() ([]models.Bank, error)
	GetByID(id uint) (*models.Bank, error)
	Create(data *models.Bank) error
	Update(id uint, data *models.Bank) error
	Delete(id uint) error
}

type bankService struct {
	repo IBankRepository
}

func NewBankService(repo IBankRepository) *bankService {
	return &bankService{repo: repo}
}

func (s *bankService) GetAll() ([]models.Bank, error) {
	return s.repo.FindAll()
}

func (s *bankService) GetByID(id uint) (*models.Bank, error) {
	return s.repo.FindByID(id)
}

func (s *bankService) Create(data *models.Bank) error {
	data.NoBank = utils.GenerateID(config.DB, "BNK", true)
	return s.repo.Create(data)
}

func (s *bankService) Update(id uint, data *models.Bank) error {
	bank, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	data.NoBank = bank.NoBank
	data.ID = bank.ID
	return s.repo.Update(data)
}

func (s *bankService) Delete(id uint) error {
	return s.repo.Delete(id)
}
