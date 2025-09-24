package utils

import (
	"errors"
	"rentroom/models"

	"gorm.io/gorm"
)

func GetProperty(db *gorm.DB, propertyID int) (models.Property, error) {
	var property models.Property
	err := db.First(&property, propertyID).Error
	if err != nil {
		return property, errors.New("property not found")
	}
	return property, nil
}

func GetUser(db *gorm.DB, userID int) (models.User, error) {
	var user models.User
	err := db.First(&user, userID).Error
	if err != nil {
		return user, errors.New("user not found")
	}
	return user, nil
}
