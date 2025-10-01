package models

type Image struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	PropertyID uint   `json:"property_id"`
	Path       string `json:"path"`
}

type ImageDeleteRequest struct {
	ImageIds []uint `json:"images_id" validate:"required"`
}
