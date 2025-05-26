package membership

import (
	"backend/models"

	"gorm.io/gorm"
)

type IMembershipRepository interface {
	FindAll() ([]models.Membership, error)
	FindByID(id uint) (*models.Membership, error)
	FindByPasienID(pasienID uint) ([]models.Membership, error)
	Create(membership *models.Membership) error
	Update(membership *models.Membership) error
	Delete(id uint) error
}

type membershipRepository struct {
	db *gorm.DB
}

func NewMembershipRepository(db *gorm.DB) IMembershipRepository {
	return &membershipRepository{db: db}
}

func (r *membershipRepository) FindAll() ([]models.Membership, error) {
	var memberships []models.Membership
	err := r.db.Preload("Pasien").Find(&memberships).Error
	return memberships, err
}

func (r *membershipRepository) FindByID(id uint) (*models.Membership, error) {
	var membership models.Membership
	err := r.db.Preload("Pasien").First(&membership, id).Error
	return &membership, err
}

func (r *membershipRepository) FindByPasienID(pasienID uint) ([]models.Membership, error) {
	var memberships []models.Membership
	err := r.db.Preload("Pasien").Where("pasien_id = ?", pasienID).Find(&memberships).Error
	return memberships, err
}

func (r *membershipRepository) Create(membership *models.Membership) error {
	return r.db.Create(membership).Error
}

func (r *membershipRepository) Update(membership *models.Membership) error {
	return r.db.Model(&models.Membership{}).Where("id = ?", membership.ID).Updates(membership).Error
}

func (r *membershipRepository) Delete(id uint) error {
	return r.db.Delete(&models.Membership{}, id).Error
}
