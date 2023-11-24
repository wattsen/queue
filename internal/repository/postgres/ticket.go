package postgres

import (
	"banking-system/internal/models"
	"gorm.io/gorm"
)

type TicketPostgres struct {
	db *gorm.DB
}

func NewTicketPostgres(db *gorm.DB) *TicketPostgres {
	return &TicketPostgres{db: db}
}

func (r *TicketPostgres) CreateTicket(ticket *models.Ticket) error {
	if err := r.db.Create(&ticket).Error; err != nil {
		return err
	}
	return nil
}

func (r *TicketPostgres) GetAllTickets() (*[]models.Ticket, error) {
	var tickets []models.Ticket
	if err := r.db.Find(&tickets).Error; err != nil {
		return nil, err
	}
	return &tickets, nil
}
