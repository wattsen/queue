package models

type Client struct {
	Id     int    `gorm:"primary_key"`
	Status string `gorm:"not null"`
	Phone  string `gorm:"not null"`
	Name   string `gorm:"not null"`
}
