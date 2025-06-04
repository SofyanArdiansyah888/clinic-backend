package stokOpname

import (
	"backend/config"
	"backend/models"
	"backend/utils"
)

type IStokOpnameService interface {
	GetAll() ([]models.StokOpname, error)
	GetByID(id uint) (*models.StokOpname, error)
	Create(data *models.StokOpname) error
}

type stokOpnameService struct {
	repo IStokOpnameRepository
}

func NewStokOpnameService(repo IStokOpnameRepository) *stokOpnameService {
	return &stokOpnameService{repo: repo}
}

func (s *stokOpnameService) GetAll() ([]models.StokOpname, error) {
	return s.repo.FindAll()
}

func (s *stokOpnameService) GetByID(id uint) (*models.StokOpname, error) {
	return s.repo.FindByID(id)
}

func (s *stokOpnameService) Create(data *models.StokOpname) error {
	data.NoStokOpname = utils.GenerateID(config.DB, "STO", true)
	return s.repo.CreateAndUpdateStok(data)
}
