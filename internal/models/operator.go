package models

type Operator struct {
	Id        uint    `gorm:"primaryKey"`
	Name      string  `gorm:"not null"`
	ServiceId uint    // Foreign key for Service, assuming it references the Service table
	Service   Service `gorm:"foreignKey:ServiceId"` // Relationship with Service
}
