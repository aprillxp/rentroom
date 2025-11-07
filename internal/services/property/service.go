package service

import (
	"errors"
	"rentroom/internal/models"
	repository "rentroom/internal/repositories/property"
	"strconv"

	"gorm.io/gorm"
)

type PropertyService struct {
	repo repository.PropertyRepository
}

func NewPropertyService(repo repository.PropertyRepository) *PropertyService {
	return &PropertyService{repo: repo}
}

func (s *PropertyService) ListPublicProperties(countryStr string) ([]models.PropertyResponse, error) {
	var countryID uint
	if countryStr != "" {
		parsed, err := strconv.ParseUint(countryStr, 10, 64)
		if err != nil {
			return nil, errors.New("invalid country id")
		}
		countryID = uint(parsed)
	}
	properties, err := s.repo.GetPublishedProperties(countryID)
	if err != nil {
		return nil, err
	}
	return NewPropertiesResponse(properties), err
}

func (s *PropertyService) GetPublishedByID(id uint) (*models.PropertyResponse, error) {
	property, err := s.repo.FindByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("property not found")
	}
	if err != nil {
		return nil, err
	}
	if property.Status != models.StatusPublished {
		return nil, errors.New("property is not published")
	}
	return NewPropertyResponse(property), err
}
