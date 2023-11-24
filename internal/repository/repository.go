package repository

import (
	"banking-system/internal/models"
	pg "banking-system/internal/repository/postgres"
	"gorm.io/gorm"
)

type Ticket interface {
	CreateTicket(ticket *models.Ticket) error
	GetAllTickets() (*[]models.Ticket, error)
}

type Appointment interface {
	CreateAppointment(appointment *models.Appointment) error
	GetAppointmentById(id uint) (*models.Appointment, error)
}

type Client interface {
	GetClientById(id uint) (*models.Client, error)
}

type Service interface {
	GetServiceById(id uint) (*models.Service, error)
}

type Repository struct {
	Ticket
	Appointment
	Service
	Client
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Ticket:      pg.NewTicketPostgres(db),
		Appointment: pg.NewAppointmentPostgres(db),
		Service:     pg.NewServicePostgres(db),
		Client:      pg.NewClientPostgres(db),
	}
}
