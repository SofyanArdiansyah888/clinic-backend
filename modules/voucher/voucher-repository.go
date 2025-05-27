package voucher

import (
	"backend/models"

	"gorm.io/gorm"
)

type IVoucherRepository interface {
	FindAll() ([]models.Voucher, error)
	FindByID(id uint) (*models.Voucher, error)
	FindByKode(kode string) (*models.Voucher, error)
	Create(voucher *models.Voucher) error
	Update(voucher *models.Voucher) error
	Delete(voucher *models.Voucher) error
}

type voucherRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) IVoucherRepository {
	return &voucherRepository{db}
}

func (r *voucherRepository) FindAll() ([]models.Voucher, error) {
	var vouchers []models.Voucher
	err := r.db.Find(&vouchers).Error
	return vouchers, err
}

func (r *voucherRepository) FindByID(id uint) (*models.Voucher, error) {
	var voucher models.Voucher
	err := r.db.First(&voucher, id).Error
	return &voucher, err
}

func (r *voucherRepository) FindByKode(kode string) (*models.Voucher, error) {
	var voucher models.Voucher
	err := r.db.Where("kode_voucher = ?", kode).First(&voucher).Error
	return &voucher, err
}

func (r *voucherRepository) Create(voucher *models.Voucher) error {
	return r.db.Create(voucher).Error
}

func (r *voucherRepository) Update(voucher *models.Voucher) error {
	return r.db.Save(voucher).Error
}

func (r *voucherRepository) Delete(voucher *models.Voucher) error {
	return r.db.Delete(voucher).Error
}
