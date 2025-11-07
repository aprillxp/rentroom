package repository

import (
	"rentroom/internal/models"
)

type PropertyRepository interface {
	GetPublishedProperties(countryID uint) ([]models.Property, error)
	FindByID(id uint) (*models.Property, error)
}
