package service

import (
	"errors"
	"rentroom/internal/models"

	"gorm.io/gorm"
)

func GetCountry(db *gorm.DB, countryID int) (models.Country, error) {
	var country models.Country
	err := db.First(&country, countryID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Country{}, errors.New("country not found")
		}
		return models.Country{}, errors.New("country not found")
	}
	return country, err
}

