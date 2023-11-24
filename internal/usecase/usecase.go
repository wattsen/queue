package usecase

import (
	"banking-system/internal/models"
	"banking-system/internal/repository"
)

type Ticket interface {
	CreateTicket(ticket *models.Ticket) (string, error)
	UpdateCoefficient() error
}

type Appointment interface {
	CreateAppointment(appointment *models.Appointment) error
	GetAppointmentById(id uint) (*models.Appointment, error)
}

type UseCase struct {
	Ticket
	Appointment
}

func NewUseCase(repo *repository.Repository) *UseCase {
	return &UseCase{
		Ticket:      NewTicketUseCase(repo.Ticket, repo.Service, repo.Client),
		Appointment: NewAppointmentUseCase(repo),
	}
}
