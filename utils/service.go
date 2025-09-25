package utils

import (
	"errors"
	"rentroom/models"

	"gorm.io/gorm"
)

func GetProperty(db *gorm.DB, propertyID int) (models.Property, error) {
	var property models.Property
	err := db.First(&property, propertyID).Error
	if err != nil {
		return property, errors.New("property not found")
	}
	return property, nil
}

func GetUser(db *gorm.DB, userID int) (models.UserResponse, error) {
	var user models.User
	err := db.First(&user, userID).Error
	if err != nil {
		return models.UserResponse{}, errors.New("user not found")
	}
	return models.UserResponse{
		ID:         user.ID,
		Username:   user.Username,
		Email:      user.Email,
		Phone:      user.Phone,
		BankID:     user.BankID,
		BankNumber: user.BankNumber,
		IsTenant:   user.IsTenant,
	}, nil
}

func GetVoucher(db *gorm.DB, voucherID int) float64 {
	var voucher models.Voucher
	err := db.First(&voucher, voucherID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0.0
		}
		return 0.0
	}
	return float64(voucher.Discount)
}

func GetTransaction(db *gorm.DB, transactionID int) (models.TransactionResponse, error) {
	var transaction models.TransactionResponse
	err := db.First(&transaction, transactionID).Error
	if err != nil {
		return models.TransactionResponse{}, errors.New("transaction not found")
	}
	return transaction, nil
}
