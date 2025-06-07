package konversiBarang

import (
	customtypes "backend/models/customTypes"
)

type CreateKonversiRequest struct {
	NoKonversi string                        `json:"no_konversi"`
	Tanggal    customtypes.DateOnly          `json:"tanggal_jt"`
	Details    []CreateKonversiDetailRequest `json:"details" validate:"required,min=1"`
}
