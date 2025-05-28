package voucher

import (
	"errors"
	"fmt"
	"time"

	"backend/config"
	"backend/models"
	"backend/utils"
)

type IVoucherService interface {
	GetAll() ([]models.Voucher, error)
	GetByID(id uint) (*models.Voucher, error)
	Create(voucher *models.Voucher) error
	Update(id uint, voucher *models.Voucher) error
	Delete(id uint) error
}

type voucherService struct {
	repo IVoucherRepository
}

func NewService(repo IVoucherRepository) IVoucherService {
	return &voucherService{repo}
}

func (s *voucherService) GetAll() ([]models.Voucher, error) {
	return s.repo.FindAll()
}

func (s *voucherService) GetByID(id uint) (*models.Voucher, error) {
	return s.repo.FindByID(id)
}

func (s *voucherService) Create(voucher *models.Voucher) error {
	// Validasi field yang wajib diisi
	if voucher.NamaVoucher == "" {
		return errors.New("nama voucher wajib diisi")
	}

	if voucher.TipeDiskon == "" {
		return errors.New("tipe diskon wajib diisi")
	}

	// Validasi tipe diskon
	if voucher.TipeDiskon != "percentage" && voucher.TipeDiskon != "fixed" {
		return errors.New("tipe diskon harus berupa 'percentage' atau 'fixed'")
	}

	// Validasi nilai diskon
	if voucher.NilaiDiskon <= 0 {
		return errors.New("nilai diskon harus lebih besar dari 0")
	}

	if voucher.TipeDiskon == "percentage" && voucher.NilaiDiskon > 100 {
		return errors.New("nilai diskon percentage tidak boleh lebih dari 100")
	}

	// Validasi minimum order dan maksimum diskon
	if voucher.MinimumOrder < 0 {
		return errors.New("minimum order tidak boleh kurang dari 0")
	}

	if voucher.MaksimumDiskon < 0 {
		return errors.New("maksimum diskon tidak boleh kurang dari 0")
	}

	// Validasi tanggal
	if voucher.TanggalMulai.IsZero() {
		voucher.TanggalMulai = time.Now()
	}

	if voucher.TanggalBerakhir.IsZero() {
		voucher.TanggalBerakhir = voucher.TanggalMulai.AddDate(0, 1, 0) // Default 1 bulan
	}

	if voucher.TanggalMulai.After(voucher.TanggalBerakhir) {
		return errors.New("tanggal mulai tidak boleh lebih besar dari tanggal berakhir")
	}

	// Generate kode voucher jika tidak diisi
	if voucher.KodeVoucher == "" {
		voucher.KodeVoucher = utils.GenerateID(config.DB, "voucher", true)
	} else {
		// Cek duplikasi kode voucher
		existing, _ := s.repo.FindByKode(voucher.KodeVoucher)
		if existing != nil {
			return errors.New("kode voucher sudah digunakan")
		}
	}

	// Set status default
	if voucher.Status == "" {
		voucher.Status = "active"
	}

	return s.repo.Create(voucher)
}

func (s *voucherService) Update(id uint, voucher *models.Voucher) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("voucher tidak ditemukan: %v", err)
	}

	// Validasi field yang wajib diisi
	if voucher.NamaVoucher == "" {
		return errors.New("nama voucher wajib diisi")
	}

	if voucher.TipeDiskon == "" {
		return errors.New("tipe diskon wajib diisi")
	}

	// Validasi tipe diskon
	if voucher.TipeDiskon != "percentage" && voucher.TipeDiskon != "fixed" {
		return errors.New("tipe diskon harus berupa 'percentage' atau 'fixed'")
	}

	// Validasi nilai diskon
	if voucher.NilaiDiskon <= 0 {
		return errors.New("nilai diskon harus lebih besar dari 0")
	}

	if voucher.TipeDiskon == "percentage" && voucher.NilaiDiskon > 100 {
		return errors.New("nilai diskon percentage tidak boleh lebih dari 100")
	}

	// Validasi minimum order dan maksimum diskon
	if voucher.MinimumOrder < 0 {
		return errors.New("minimum order tidak boleh kurang dari 0")
	}

	if voucher.MaksimumDiskon < 0 {
		return errors.New("maksimum diskon tidak boleh kurang dari 0")
	}

	// Validasi tanggal
	if !voucher.TanggalMulai.IsZero() && !voucher.TanggalBerakhir.IsZero() {
		if voucher.TanggalMulai.After(voucher.TanggalBerakhir) {
			return errors.New("tanggal mulai tidak boleh lebih besar dari tanggal berakhir")
		}
	}

	// Cek perubahan kode voucher
	if voucher.KodeVoucher != "" && voucher.KodeVoucher != existing.KodeVoucher {
		existingKode, _ := s.repo.FindByKode(voucher.KodeVoucher)
		if existingKode != nil {
			return errors.New("kode voucher sudah digunakan")
		}
	}

	// Update fields
	existing.NamaVoucher = voucher.NamaVoucher
	existing.Deskripsi = voucher.Deskripsi
	existing.TipeDiskon = voucher.TipeDiskon
	existing.NilaiDiskon = voucher.NilaiDiskon
	existing.MinimumOrder = voucher.MinimumOrder
	existing.MaksimumDiskon = voucher.MaksimumDiskon
	existing.Kuota = voucher.Kuota

	if !voucher.TanggalMulai.IsZero() {
		existing.TanggalMulai = voucher.TanggalMulai
	}
	if !voucher.TanggalBerakhir.IsZero() {
		existing.TanggalBerakhir = voucher.TanggalBerakhir
	}
	if voucher.Status != "" {
		existing.Status = voucher.Status
	}
	if voucher.KodeVoucher != "" {
		existing.KodeVoucher = voucher.KodeVoucher
	}

	return s.repo.Update(existing)
}

func (s *voucherService) Delete(id uint) error {
	existing, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("voucher tidak ditemukan: %v", err)
	}

	return s.repo.Delete(existing)
}
