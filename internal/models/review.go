package models

type Review struct {
	ID            uint    `gorm:"primaryKey" json:"id"`
	PropertyID    uint    `json:"property_id"`
	TransactionID uint    `json:"transaction_id"`
	Rating        float64 `json:"rating"`
	Description   string  `json:"description"`
}

type ReviewRequest struct {
	Rating      float64 `json:"rating" validate:"required,gte=0,lte=5"`
	Description string  `json:"description" validate:"omitempty,min=5"`
}
