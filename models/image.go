package models

type Image struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	PropertyID uint   `gorm:"not null;index"`
	Path       string `json:"path"`

	Property Property `gorm:"foreignKey:PropertyID"`
}

type ImageDeleteRequest struct {
	ImageIds []uint `json:"images_id" validate:"required"`
}

type ImageResponse struct {
	ID         uint   `json:"id"`
	PropertyID uint   `json:"property_id"`
	Path       string `json:"path"`
}
