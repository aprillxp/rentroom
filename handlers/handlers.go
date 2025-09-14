package handlers

import (
	"net/http"
	auth "rentroom/handlers/auth"

	"gorm.io/gorm"
)

func UserRegister(db *gorm.DB) http.HandlerFunc {
	return auth.UserRegister(db)
}

func UserLogin(db *gorm.DB) http.HandlerFunc {
	return auth.UserLogin(db)
}
