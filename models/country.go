package models

type Country struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Name        string  `json:"name"`
	Path        *string `json:"banner"`
	Description *string `json:"description"`
}
type CountryRequest struct {
	Name        string  `json:"name" validate:"required,min=3,max=50"`
	Description *string `json:"description,omitempty" validate:"omitempty,min=10"`
}
