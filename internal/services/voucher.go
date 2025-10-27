package service

import (
	"errors"
	"rentroom/internal/models"

	"gorm.io/gorm"
)

func GetVoucher(db *gorm.DB, voucherID int) (models.Voucher, error) {
	var voucher models.Voucher
	err := db.First(&voucher, voucherID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Voucher{}, errors.New("voucher not found")
		}
		return models.Voucher{}, errors.New("voucher not found")
	}
	return voucher, err
}