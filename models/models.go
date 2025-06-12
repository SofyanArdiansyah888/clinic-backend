package models

type Models []interface{}

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
