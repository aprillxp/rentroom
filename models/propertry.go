package models

import (
	"time"
)

const (
	StatusPublished = 1
	StatusDraft     = 2
)

type Property struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	Name             string    `json:"name"`
	CountryID        uint      `json:"country_id"`
	Guests           int       `json:"guests"`
	Price            float64   `json:"price"`
	Status           int       `json:"status"`
	DisabledDateFrom time.Time `gorm:"type:date" json:"disabled_date_from"`
	DisabledDateTo   time.Time `gorm:"type:date" json:"disabled_date_to"`
	Description      string    `json:"description"`
	Geo              string    `json:"geo"`
	Province         string    `json:"province"`
	District         string    `json:"district"`
	City             string    `json:"city"`
	Address          string    `json:"address"`
	Zip              string    `json:"zip"`

	Images []Image `gorm:"constraint:OnDelete:CASCADE"`
}

type PropertyAmenities struct {
	PropertyID uint `gorm:"primaryKey;not null"`
	AmenityID  uint `gorm:"primaryKey;not null"`

	Property Property `gorm:"foreignKey:PropertyID;references:ID;constraint:OnDelete:CASCADE"`
	Amenity  Amenity  `gorm:"foreignKey:AmenityID;references:ID;constraint:OnDelete:CASCADE"`
}

type PropertyCreateRequest struct {
	Name             string    `json:"name" validate:"required,min=3"`
	CountryID        uint      `json:"country_id" validate:"required,gt=0"`
	Guests           int       `json:"guests" validate:"required,gt=0"`
	Price            float64   `json:"price" validate:"required,gt=0"`
	DisabledDateFrom time.Time `json:"disabled_date_from" validate:"required"`
	DisabledDateTo   time.Time `json:"disabled_date_to" validate:"required"`
	Description      string    `json:"description" validate:"required,min=10"`
	Geo              string    `json:"geo" validate:"required,min=3"`
	Province         string    `json:"province" validate:"required,min=2"`
	District         string    `json:"district" validate:"required,min=2"`
	City             string    `json:"city" validate:"required,min=2"`
	Address          string    `json:"address" validate:"required,min=5"`
	Zip              string    `json:"zip" validate:"required,gt=0"`
	Amenities        []uint    `json:"amenities" validate:"required,min=1"`
}
type PropertyEditRequest struct {
	Name             *string    `json:"name" validate:"omitempty,min=3"`
	CountryID        *uint      `json:"country_id" validate:"omitempty,gt=0"`
	Guests           *int       `json:"guests" validate:"omitempty,gt=0"`
	Price            *float64   `json:"price" validate:"omitempty,gt=0"`
	DisabledDateFrom *time.Time `json:"disabled_date_from" validate:"omitempty"`
	DisabledDateTo   *time.Time `json:"disabled_date_to" validate:"omitempty"`
	Description      *string    `json:"description" validate:"omitempty,min=10"`
	Geo              *string    `json:"geo" validate:"omitempty,min=3"`
	Province         *string    `json:"province" validate:"omitempty,min=2"`
	District         *string    `json:"district" validate:"omitempty,min=2"`
	City             *string    `json:"city" validate:"omitempty,min=2"`
	Address          *string    `json:"address" validate:"omitempty,min=5"`
	Zip              *string    `json:"zip" validate:"omitempty,gt=0"`
	Amenities        *[]uint    `json:"amenities" validate:"omitempty,min=1"`
}

type PropertyResponse struct {
	ID               uint      `json:"id"`
	Name             string    `json:"name"`
	CountryID        uint      `json:"country_id"`
	Guests           int       `json:"guests"`
	Price            float64   `json:"price"`
	Status           int       `json:"status"`
	DisabledDateFrom time.Time `gorm:"type:date" json:"disabled_date_from"`
	DisabledDateTo   time.Time `gorm:"type:date" json:"disabled_date_to"`
	Description      string    `json:"description"`
	Geo              string    `json:"geo"`
	Province         string    `json:"province"`
	District         string    `json:"district"`
	City             string    `json:"city"`
	Address          string    `json:"address"`
	Zip              string    `json:"zip"`
}
