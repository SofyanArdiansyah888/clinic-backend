package appointment

import (
	"backend/models"

	"gorm.io/gorm"
)

type IAppointmentRepository interface {
	FindAll() ([]models.Appointment, error)
	FindByID(id uint) (*models.Appointment, error)
	Create(appointment *models.Appointment) error
	Update(appointment *models.Appointment) error
	Delete(id uint) error
}

type appointmentRepository struct {
	db *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) IAppointmentRepository {
	return &appointmentRepository{db: db}
}

func (r *appointmentRepository) FindAll() ([]models.Appointment, error) {
	var appointments []models.Appointment
	err := r.db.Preload("Pasien").Preload("Staff").Preload("Cabang").Find(&appointments).Error
	return appointments, err
}

func (r *appointmentRepository) FindByID(id uint) (*models.Appointment, error) {
	var appointment models.Appointment
	err := r.db.Preload("Pasien").Preload("Staff").Preload("Cabang").First(&appointment, id).Error
	return &appointment, err
}

func (r *appointmentRepository) Create(appointment *models.Appointment) error {
	return r.db.Create(appointment).Error
}

func (r *appointmentRepository) Update(appointment *models.Appointment) error {
	return r.db.Model(&models.Appointment{}).Where("id = ?", appointment.ID).Updates(appointment).Error
}

func (r *appointmentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Appointment{}, id).Error
}
