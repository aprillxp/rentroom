package service

import "rentroom/internal/models"

func NewPropertiesResponse(properties []models.Property) []models.PropertyResponse {
	responses := make([]models.PropertyResponse, 0, len(properties))
	for _, p := range properties {
		responses = append(responses, models.PropertyResponse{
			ID:               p.ID,
			CountryID:        p.CountryID,
			Name:             p.Name,
			Guests:           p.Guests,
			Price:            p.Price,
			Status:           p.Status,
			DisabledDateFrom: p.DisabledDateFrom,
			DisabledDateTo:   p.DisabledDateTo,
			Description:      p.Description,
			Geo:              p.Geo,
			Province:         p.Province,
			District:         p.District,
			City:             p.City,
			Address:          p.Address,
			Zip:              p.Zip,
		})
	}
	return responses
}

func NewPropertyResponse(p *models.Property) *models.PropertyResponse {
	return &models.PropertyResponse{
		ID:               p.ID,
		CountryID:        p.CountryID,
		Name:             p.Name,
		Guests:           p.Guests,
		Price:            p.Price,
		Status:           p.Status,
		DisabledDateFrom: p.DisabledDateFrom,
		DisabledDateTo:   p.DisabledDateTo,
		Description:      p.Description,
		Geo:              p.Geo,
		Province:         p.Province,
		District:         p.District,
		City:             p.City,
		Address:          p.Address,
		Zip:              p.Zip,
	}
}
