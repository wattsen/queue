package delivery

import (
	"banking-system/internal/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase *usecase.UseCase
}

func NewHandler(useCase *usecase.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.New()

	tickets := router.Group("/tickets")
	{
		tickets.POST("/", h.CreateTicket)

	}

	appointments := router.Group("/appointments")
	{
		appointments.POST("/", h.CreateAppointment)
		//appointment.GET("/:id", h.GetAppointmentByID)

	}
	return router
}
