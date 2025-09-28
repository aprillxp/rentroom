package models

import "time"

const (
	StatusApproved = 1
	StatusPending  = 2
	StatusCanceled = 3
	StatusRejected = 4
	StatusDone     = 5
)

type Transaction struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	UserID     uint      `json:"user_id"`
	PropertyID uint      `json:"property_id"`
	Price      float64   `json:"price"`
	CheckIn    time.Time `gorm:"type:date" json:"check_in"`
	CheckOut   time.Time `gorm:"type:date" json:"check_out"`
	Status     int       `json:"status"`
	VoucherID  uint      `json:"voucher_id"`
}

type TransactionRequest struct {
	PropertyID uint      `json:"property_id" validate:"required"`
	CheckIn    time.Time `json:"check_in" validate:"required"`
	CheckOut   time.Time `json:"check_out" validate:"required,gtfield=CheckIn"`
	VoucherID  *uint     `json:"voucher_id" validate:"omitempty"`
}

type TransactionResponse struct {
	ID         uint      `json:"id"`
	PropertyID uint      `json:"property_id"`
	Price      float64   `json:"price"`
	CheckIn    time.Time `gorm:"type:date" json:"check_in"`
	CheckOut   time.Time `gorm:"type:date" json:"check_out"`
	Status     int       `json:"status"`
	VoucherID  uint      `json:"voucher_id"`
}
