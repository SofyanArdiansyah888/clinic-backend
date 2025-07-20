package antrian

import (
	"backend/models"
	"fmt"
)

type IAntrianService interface {
	GetAll() ([]models.Antrian, error)
	GetByID(id uint) (*models.Antrian, error)
	createPasienAntrian(pasien *models.Pasien, antrian *models.Antrian) error
	Create(data *models.Antrian) error
	Update(id uint, data *models.Antrian) error
	Delete(id uint) error
}

type antrianService struct {
	repo IAntrianRepository
}

func NewAntrianService(repo IAntrianRepository) *antrianService {
	return &antrianService{repo: repo}
}

func (s *antrianService) GetAll() ([]models.Antrian, error) {
	return s.repo.FindAll()
}

func (s *antrianService) GetByID(id uint) (*models.Antrian, error) {
	return s.repo.FindByID(id)
}

func (s *antrianService) Create(data *models.Antrian) error {
	return s.repo.Create(data)
}
func (s *antrianService) createPasienAntrian(pasien *models.Pasien, antrian *models.Antrian) error {
	// Validate input data before processing
	if pasien == nil || antrian == nil {
		return fmt.Errorf("data pasien atau antrian tidak boleh kosong")
	}
	if pasien.NamaPasien == "" {
		return fmt.Errorf("nama pasien wajib diisi")
	}

	// Check and validate required fields
	// (Tambahkan validasi field lain sesuai kebutuhan)

	// Ambil DB dari repo (repo harus bertipe *antrianRepository)
	repoImpl, ok := s.repo.(*antrianRepository)
	if !ok || repoImpl.db == nil {
		return fmt.Errorf("repo tidak valid")
	}

	tx := repoImpl.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// 1. Buat pasien
	if err := tx.Create(pasien).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("gagal membuat pasien: %w", err)
	}

	// 2. Set pasien_id pada antrian
	antrian.IDPasien = pasien.ID

	// 3. Buat antrian
	if err := tx.Omit("Pasien", "Staff").Create(antrian).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("gagal membuat antrian: %w", err)
	}

	// Commit transaksi
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("gagal commit transaksi: %w", err)
	}

	// Log create operations for audit purposes
	// (Implementasi log sesuai kebutuhan, misal: log.Printf("Create antrian: %v", antrian))

	return nil
}

func (s *antrianService) Update(id uint, data *models.Antrian) error {
	antrian, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	data.NoAntrian = antrian.NoAntrian
	data.ID = antrian.ID
	return s.repo.Update(data)
}

func (s *antrianService) Delete(id uint) error {
	return s.repo.Delete(id)
}
