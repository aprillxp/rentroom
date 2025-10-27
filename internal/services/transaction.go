package service

import (
	"errors"
	"rentroom/internal/models"

	"gorm.io/gorm"
)

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
		UserID:     transaction.UserID,
		PropertyID: transaction.PropertyID,
		Price:      transaction.Price,
		CheckIn:    transaction.CheckIn,
		CheckOut:   transaction.CheckOut,
		Status:     transaction.Status,
		VoucherID:  transaction.VoucherID,
	}, nil
}
