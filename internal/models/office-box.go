package models

type OfficeBox struct {
	Id               uint     `gorm:"primary_key"`
	Name             string   `gorm:"not null"`
	TicketId         uint     `gorm:""`                      // Foreign key for Ticket, assuming it references the Ticket table
	Ticket           Ticket   `gorm:"foreignKey:TicketId"`   // Relationship with Ticket
	OperatorId       int      `gorm:"not null"`              // Foreign key for Operator, assuming it references the Operator table
	Operator         Operator `gorm:"foreignKey:OperatorId"` // Relationship with Operator
	StatusOfBusiness string   `gorm:""`
}
