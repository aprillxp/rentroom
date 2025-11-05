package service

import "rentroom/internal/models"

func NewUserResponse(u *models.User) *models.UserResponse {
	return &models.UserResponse{
		ID:         u.ID,
		Username:   u.Username,
		Email:      u.Email,
		Phone:      u.Phone,
		Bank:       u.Bank,
		BankNumber: u.BankNumber,
		IsTenant:   u.IsTenant,
	}
}
