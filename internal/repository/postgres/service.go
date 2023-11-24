package postgres

import (
	"banking-system/internal/models"
	"gorm.io/gorm"
)

type ServicePostgres struct {
	db *gorm.DB
}

func NewServicePostgres(db *gorm.DB) *ServicePostgres {
	return &ServicePostgres{db: db}
}

func (r *ServicePostgres) GetServiceById(id uint) (*models.Service, error) {
	var service models.Service
	if err := r.db.Where("id = ?", id).First(&service).Error; err != nil {
		return nil, err
	}
	return &service, nil
}
