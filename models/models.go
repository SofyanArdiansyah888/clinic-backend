package models

// Models is a slice of pointers to all model structs for migration/registration.
type Models []any

// GetModels returns a slice of all model pointers for GORM auto-migration.
func GetModels() Models {
	return Models{
		&MonthlySequence{},
		&Barang{},
		&StokMovement{},
		&StokOpname{},
		&LoginHistory{},
		&User{},
		&Pasien{},
		&Staff{},
		&Treatment{},
		&Antrian{},
		&Supplier{},
		&Perawatan{},
		&Cabang{},
		&Bank{},
		&TemplateConcern{},
		&Appointment{},
		&AppointmentTreatment{},
		&Membership{},
		&Voucher{},
		&Pembelian{},
		&PembelianDetail{},
		&Penjualan{},
		&PenjualanDetail{},
		&KonversiStok{},
		&KonversiStokDetail{},
		&ProduksiBarang{},
		&ProduksiBarangDetail{},
	}
}
