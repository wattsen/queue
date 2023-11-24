package delivery

import (
	"banking-system/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateAppointment(c *gin.Context) {
	var input *models.Appointment
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input JSON"})
		return
	}

	err := h.useCase.CreateAppointment(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while creating Appointment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Appointment created successfully",
	})
}

//func (h *Handler) GetAppointmentByID(c *gin.Context) {
//	id := c.Param("id")
//
//	appointment, err := h.useCase.GetAppointmentByID(id)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while getting Appointment"})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"message": "Appointment fetched successfully",
//		"data":    appointment,
//	})
//}
