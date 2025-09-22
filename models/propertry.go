package models

import (
	"time"
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
	Zip              int       `json:"zip"`
	Amenities        string    `json:"amenities"`
}

type PropertyRequest struct {
	Name             string    `json:"name" validate:"required"`
	CountryID        uint      `json:"country_id" validate:"required"`
	Guests           int       `json:"guests" validate:"required"`
	Price            float64   `json:"price" validate:"required"`
	Status           int       `json:"status" validate:"required"`
	DisabledDateFrom time.Time `gorm:"type:date" json:"disabled_date_from" validate:"required"`
	DisabledDateTo   time.Time `gorm:"type:date" json:"disabled_date_to" validate:"required"`
	Description      string    `json:"description" validate:"required"`
	Geo              string    `json:"geo" validate:"required"`
	Province         string    `json:"province" validate:"required"`
	District         string    `json:"district" validate:"required"`
	City             string    `json:"city" validate:"required"`
	Address          string    `json:"address" validate:"required"`
	Zip              int       `json:"zip" validate:"required"`
	Amenities        string    `json:"amenities" validate:"required"`
}
