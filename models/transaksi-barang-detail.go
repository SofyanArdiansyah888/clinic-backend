package models

type TransaksiBarangDetail struct {
    ID             uint    `json:"id" gorm:"primaryKey"`
    IDCabang       uint    `json:"id_cabang" gorm:"not null"`
    NomorTransaksi string  `json:"nomor_transaksi" gorm:"type:varchar(100);not null"`
    KodeBarang     string  `json:"kode_barang" gorm:"type:varchar(50);not null"`
    Harga          float64 `json:"harga" gorm:"type:decimal(15,2);not null"`
    Jumlah         int     `json:"jumlah" gorm:"not null"`
    TanggalExpired string  `json:"tanggal_expired" gorm:"type:date;default:null"`
    Diskon         float64 `json:"diskon" gorm:"type:decimal(15,2);not null;default:0"`
    PPN            float64 `json:"ppn" gorm:"type:decimal(15,2);not null;default:0"`
    HPP            float64 `json:"hpp" gorm:"type:decimal(15,2);not null;default:0"`
    Ongkir         float64 `json:"ongkir" gorm:"type:decimal(15,2);not null;default:0"`
    Tipe           string  `json:"tipe" gorm:"type:varchar(20);not null"` // keluar, masuk
    Cabang          Cabang          `json:"cabang" gorm:"foreignKey:IDCabang"`
    TransaksiBarang TransaksiBarang `json:"transaksi_barang" gorm:"foreignKey:NomorTransaksi;references:NomorTransaksi"`
    Barang          Barang          `json:"barang" gorm:"foreignKey:KodeBarang;references:KodeBarang"`
}