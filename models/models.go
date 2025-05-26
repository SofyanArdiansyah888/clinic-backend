package models

type Models []interface{}

func GetModels() Models {
	return Models{
		&MonthlySequence{},
		&Barang{},
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
	}
}
