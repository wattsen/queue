package models

type Branch struct {
	Id      uint     `gorm:"primary_key"`
	Name    string   `gorm:"not null"`
	Address string   `gorm:"not null"`
	Tickets []Ticket // Relationship definition
}
