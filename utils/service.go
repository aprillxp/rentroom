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

func GetUser(db *gorm.DB, userID uint) (models.UserResponse, error) {
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

func GetCountry(db *gorm.DB, countryID int) (models.Country, error) {
	var country models.Country
	err := db.First(&country, countryID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Country{}, errors.New("country not found")
		}
		return models.Country{}, errors.New("country not found")
	}
	return country, err
}

func GetTransaction(db *gorm.DB, transactionID uint) (models.TransactionResponse, error) {
	var transaction models.Transaction
	err := db.
		Where("id = ?", transactionID).
		First(&transaction).Error
	if err != nil {
		return models.TransactionResponse{}, errors.New("transaction not found")
	}
	return models.TransactionResponse{
		ID:         transaction.ID,
		PropertyID: transaction.PropertyID,
		Price:      transaction.Price,
		CheckIn:    transaction.CheckIn,
		CheckOut:   transaction.CheckOut,
		Status:     transaction.Status,
		VoucherID:  transaction.VoucherID,
	}, nil
}

func GetUserTransaction(db *gorm.DB, userID, transactionID uint) (models.TransactionResponse, error) {
	var transaction models.Transaction
	err := db.
		Where("id = ? AND user_id = ?", transactionID, userID).
		First(&transaction).Error
	if err != nil {
		return models.TransactionResponse{}, errors.New("transaction not found")
	}
	return models.TransactionResponse{
		ID:         transaction.ID,
		PropertyID: transaction.PropertyID,
		Price:      transaction.Price,
		CheckIn:    transaction.CheckIn,
		CheckOut:   transaction.CheckOut,
		Status:     transaction.Status,
		VoucherID:  transaction.VoucherID,
	}, nil
}

func GetTenantTransaction(db *gorm.DB, propertyIDs []uint, transactionID uint) (models.TransactionResponse, error) {
	var transaction models.Transaction
	err := db.
		Where("id = ? AND property_id IN ?", transactionID, propertyIDs).
		First(&transaction).Error
	if err != nil {
		return models.TransactionResponse{}, errors.New("transaction not found")
	}
	return models.TransactionResponse{
		ID:         transaction.ID,
		PropertyID: transaction.PropertyID,
		Price:      transaction.Price,
		CheckIn:    transaction.CheckIn,
		CheckOut:   transaction.CheckOut,
		Status:     transaction.Status,
		VoucherID:  transaction.VoucherID,
	}, nil
}

func GetPropertyIDs(db *gorm.DB, userID uint) ([]uint, error) {
	var propertyIDs []uint
	err := db.Model(&models.UserProperties{}).
		Where("user_id = ?", userID).
		Pluck("property_id", &propertyIDs).Error
	if err != nil {
		return nil, err
	}
	if len(propertyIDs) == 0 {
		return nil, errors.New("no properties found for this user")
	}
	return propertyIDs, nil
}
