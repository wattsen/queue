package postgres

import (
	"banking-system/internal/models"
	"gorm.io/gorm"
)

type AppointmentPostgres struct {
	db *gorm.DB
}

func NewAppointmentPostgres(db *gorm.DB) *AppointmentPostgres {
	return &AppointmentPostgres{db: db}
}

func (r *AppointmentPostgres) CreateAppointment(appointment *models.Appointment) error {
	if err := r.db.Create(appointment).Error; err != nil {
		return err
	}
	return nil
}

func (r *AppointmentPostgres) GetAppointmentById(id uint) (*models.Appointment, error) {
	var appointment models.Appointment
	if err := r.db.Where("id = ?", id).First(&appointment).Error; err != nil {
		return nil, err
	}
	return &appointment, nil
}
