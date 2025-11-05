package repository

import "rentroom/internal/models"

type UserRepository interface {
	FindByIdentifier(identifier string) (*models.User, error)
	Create(user *models.User) error
	Update(user *models.User) error
	FindByID(id uint) (*models.User, error)
	UpdateFields(id uint, updates map[string]interface{}) error

	ExistsUsername(username string, excludeID uint) (bool, error)
	ExistsEmail(email string, excludeID uint) (bool, error)
	ExistsPhone(phone string, excludeID uint) (bool, error)
}
