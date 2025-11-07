package repository

import (
	"rentroom/internal/models"

	"gorm.io/gorm"
)

type GormPropertyRepository struct {
	db *gorm.DB
}

func NewGormPropertyRepository(db *gorm.DB) PropertyRepository {
	return &GormPropertyRepository{db: db}
}

func (r *GormPropertyRepository) GetPublishedProperties(countryID uint) ([]models.Property, error) {
	var properties []models.Property
	query := r.db
	if countryID != 0 {
		query = query.Where("country_id = ? AND status = ?", countryID, models.StatusPublished)
	} else {
		query = query.Where("status = ?", models.StatusPublished)
	}
	err := query.Find(&properties).Error
	if err != nil {
		return nil, err
	}
	return properties, nil
}

func (r *GormPropertyRepository) FindByID(id uint) (*models.Property, error) {
	var property models.Property
	err := r.db.First(&property, id).Error
	if err != nil {
		return nil, err
	}
	return &property, nil
}
