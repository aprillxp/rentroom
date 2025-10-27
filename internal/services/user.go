package service

import (
	"errors"
	"rentroom/internal/models"

	"gorm.io/gorm"
)

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