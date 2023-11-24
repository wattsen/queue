package models

import "time"

type Ticket struct {
	Id            uint      `gorm:"primaryKey;autoIncrement"`
	TicketID      string    `gorm:"not null;unique"`
	ClientID      uint      `gorm:""`
	ServiceID     uint      `gorm:"not null"`
	Coefficient   int       `gorm:"not null"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	BranchID      uint      `gorm:"not null"` // Foreign key
	Branch        Branch    `gorm:"foreignKey:BranchID"`
	IsAppointment bool      `gorm:"not null;default:false"` // Relationship definition
}
