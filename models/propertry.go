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
	Amenities        []Amenity `gorm:"many2many:property_amenities;" json:"amenities"`
	DisabledDateFrom time.Time `gorm:"type:date" json:"disabled_date_from"`
	DisabledDateTo   time.Time `gorm:"type:date" json:"disabled_date_to"`
	Description      string    `json:"description"`
	Geo              string    `json:"geo"`
	Province         string    `json:"province"`
	District         string    `json:"district"`
	City             string    `json:"city"`
	Address          string    `json:"address"`
	Zip              int       `json:"zip"`
}
