package delivery

import (
	"banking-system/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateTicket(c *gin.Context) {
	var input *models.Ticket
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input JSON"})
		return
	}

	ticketId, err := h.useCase.CreateTicket(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ticketId": ticketId,
	})
}
