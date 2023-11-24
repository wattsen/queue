package models

import "time"

type Appointment struct {
	Id            uint      `gorm:"primary_key"`
	ClientId      uint      `gorm:"not null"`
	ScheduledTime time.Time `gorm:"not null"`
}
