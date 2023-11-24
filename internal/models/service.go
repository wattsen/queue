package models

type Service struct {
	Id          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Coefficient int    `gorm:"not null"`
}
