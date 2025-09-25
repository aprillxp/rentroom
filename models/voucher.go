package models

import "time"

type Voucher struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `gorm:"unique" json:"username"`
	Discount   int       `json:"discount"`
	Quantity   int       `json:"quantity"`
	EndPeriode time.Time `gorm:"type:date" json:"end_periode"`
}
