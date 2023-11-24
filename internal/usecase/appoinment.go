package usecase

import (
	"banking-system/internal/models"
	"banking-system/internal/repository"
)

type AppointmentUseCase struct {
	repo repository.Appointment
}

func NewAppointmentUseCase(repo repository.Appointment) *AppointmentUseCase {
	return &AppointmentUseCase{
		repo: repo,
	}
}

func (a *AppointmentUseCase) CreateAppointment(appointment *models.Appointment) error {
	return a.repo.CreateAppointment(appointment)
}

func (a *AppointmentUseCase) GetAppointmentById(id uint) (*models.Appointment, error) {
	return a.repo.GetAppointmentById(id)
}
