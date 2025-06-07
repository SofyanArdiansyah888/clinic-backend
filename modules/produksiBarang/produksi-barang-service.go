package produksiBarang

import (
	"backend/config"
	"backend/models"
	"backend/utils"
)

type ProduksiBarangService struct {
	repo *ProduksiBarangRepository
}

func NewProduksiBarangService(repo *ProduksiBarangRepository) *ProduksiBarangService {
	return &ProduksiBarangService{repo: repo}
}

func (s *ProduksiBarangService) Create(req CreateProduksiRequest) (*models.ProduksiBarang, error) {

	// Generate nomor transaksi
	nomorTransaksi := utils.GenerateID(config.DB, "PRD", true)

	// Prepare transaksi header
	transaksi := &models.ProduksiBarang{
		NoProduksi: nomorTransaksi,
		Tanggal:    req.Tanggal,
	}

	// Prepare detail transaksi
	details := make([]models.ProduksiBarangDetail, len(req.Details))

	for i, d := range req.Details {
		details[i] = models.ProduksiBarangDetail{
			KodeBarang: d.KodeBarang,
			NoProduksi: d.NoProduksi,
			Quantity:   d.Quantity,
			Arah:       d.Arah,
		}
	}

	// Execute transaction
	err := s.repo.CreateTransaksiAndUpdateStock(transaksi, details)
	if err != nil {
		return nil, err
	}

	return transaksi, nil
}

func (s *ProduksiBarangService) GetByNomor(nomorTransaksi string) (*models.ProduksiBarang, []models.ProduksiBarangDetail, error) {
	return s.repo.FindByNomor(nomorTransaksi)
}
