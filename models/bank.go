package models

type Bank struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}
