package kartuStok

import (
	"backend/models"
)

type IKartuStokService interface {
	GetKartuStok(kodeBarang string, dari string, sampai string) ([]models.StokMovement, error)
}

type kartuStokService struct {
	repo IKartuStokRepository
}

func NewKartuStokService(repo IKartuStokRepository) *kartuStokService {
	return &kartuStokService{repo: repo}
}

func (s *kartuStokService) GetKartuStok(kodeBarang string, dari string, sampai string) ([]models.StokMovement, error) {
	result, err := s.repo.GetStokMovement(kodeBarang, dari, sampai)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, nil
	}
	return result, nil
}
