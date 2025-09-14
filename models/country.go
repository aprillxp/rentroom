package models

type Country struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Banner      string `json:"banner"`
	Description string `json:"description"`
}
