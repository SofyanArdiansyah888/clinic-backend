package produksiBarang

import (
	customtypes "backend/models/customTypes"
)

type CreateProduksiRequest struct {
	NoProduksi string                        `json:"no_produksi"`
	Tanggal    customtypes.DateOnly          `json:"tanggal"`
	Details    []CreateProduksiDetailRequest `json:"details" validate:"required,min=1"`
}
