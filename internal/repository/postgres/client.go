package postgres

import (
	"banking-system/internal/models"
	"gorm.io/gorm"
)

type ClientPostgres struct {
	db *gorm.DB
}

func NewClientPostgres(db *gorm.DB) *ClientPostgres {
	return &ClientPostgres{db: db}
}

func (c *ClientPostgres) GetClientById(id uint) (*models.Client, error) {
	var client models.Client
	err := c.db.First(&client, id).Error
	if err != nil {
		return nil, err
	}
	return &client, nil
}
