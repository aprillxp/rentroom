package models

import "time"

type Voucher struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `gorm:"unique" json:"name"`
	Discount   float64   `json:"discount"`
	Quantity   int       `json:"quantity"`
	EndPeriode time.Time `gorm:"type:date" json:"end_periode"`
}

type VoucherRequest struct {
	Name       string    `json:"name" validate:"required,min=3,max=50"`
	Discount   float64   `json:"discount" validate:"required,gte=0,lte=1"`
	Quantity   int       `json:"quantity" validate:"required,gte=1"`
	EndPeriode time.Time `json:"end_periode" validate:"required"`
}
