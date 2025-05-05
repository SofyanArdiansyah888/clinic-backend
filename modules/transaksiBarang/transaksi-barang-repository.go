package transaksiBarang

import "backend/models"

type ITransaksiBarangRepository interface {
	CreateTransaksi() (models.Barang, error)
}
type transaksiBarangRepository struct{}

func (t transaksiBarangRepository) CreateTransaksi() (models.Barang, error) {
	//TODO implement me
	panic("implement me")
}

func NewTransaksiBarangRepository() ITransaksiBarangRepository {
	return &transaksiBarangRepository{}
}
