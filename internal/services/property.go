package service

import (
	"errors"
	"rentroom/internal/models"

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


func GetPropertyIDs(db *gorm.DB, userID uint) ([]uint, error) {
	var propertyIDs []uint
	err := db.Model(&models.UserProperties{}).
		Where("user_id = ?", userID).
		Pluck("property_id", &propertyIDs).Error
	if err != nil {
		return nil, err
	}
	if len(propertyIDs) == 0 {
		return nil, errors.New("no properties found for this user")
	}
	return propertyIDs, nil
}
