package usecase

import (
	"banking-system/internal/models"
	"banking-system/internal/repository"
	"fmt"
	"math/rand"
	"time"
)

var coeffForClient = map[string]int{"Gold": 10, "Platinum": 20, "VIP": 30}
var coeffForTime = 20

type TicketUseCase struct {
	repoTicket  repository.Ticket
	repoService repository.Service
	repoClient  repository.Client
}

func NewTicketUseCase(repoTicket repository.Ticket, repoService repository.Service, repoClient repository.Client) *TicketUseCase {
	return &TicketUseCase{
		repoTicket:  repoTicket,
		repoService: repoService,
		repoClient:  repoClient,
	}
}

func (uc *TicketUseCase) CreateTicket(ticket *models.Ticket) (string, error) {
	// todo проверить нынешнее время
	currentTime := time.Now()
	startTime := currentTime.Add(-5 * time.Minute)
	endTime := currentTime.Add(5 * time.Minute)
	if ticket.IsAppointment && (currentTime.After(startTime) && currentTime.Before(endTime)) {
		ticket.Coefficient += coeffForTime
	}

	service, err := uc.repoService.GetServiceById(ticket.ServiceID)
	if err != nil {
		return "", err
	}
	ticket.Coefficient += service.Coefficient

	ticketId := generateRandomString()
	ticket.TicketID = ticketId

	if ticket.ClientID != 0 {
		client, err := uc.repoClient.GetClientById(ticket.ClientID)
		if err != nil {
			return "", err
		}
		ticket.Coefficient += coeffForClient[client.Status]
	}

	err = uc.repoTicket.CreateTicket(ticket)
	if err != nil {
		return "", err
	}

	return ticketId, nil
}

func generateRandomString() string {
	rand.Seed(time.Now().UnixNano())

	// Generate random letters
	letter1 := byte(rand.Intn(26) + 'A')
	letter2 := byte(rand.Intn(26) + 'A')

	// Generate random digits
	digit1 := rand.Intn(10)
	digit2 := rand.Intn(10)
	digit3 := rand.Intn(10)

	// Create the string with the specified format
	result := fmt.Sprintf("%c%c%d%d%d", letter1, letter2, digit1, digit2, digit3)
	return result
}

func (uc *TicketUseCase) UpdateCoefficient() error {
	currentTime := time.Now()
	startTime := currentTime.Add(-5 * time.Minute)
	endTime := currentTime.Add(5 * time.Minute)

	tickets, err := uc.repoTicket.GetAllTickets()
	if err != nil {
		return err
	}

	for _, ticket := range *tickets {
		if ticket.IsAppointment && !(currentTime.After(startTime) && currentTime.Before(endTime)) {
			serviceCoeff, err := uc.repoService.GetServiceById(ticket.ServiceID)
			if err != nil {
				return err
			}
			ticket.Coefficient = serviceCoeff.Coefficient
		}
	}

	return nil
}
