package handlers

import (
	"net/http"
	auth "rentroom/handlers/auth"
	property "rentroom/handlers/properties"

	"gorm.io/gorm"
)

func UserRegister(db *gorm.DB) http.HandlerFunc {
	return auth.UserRegister(db)
}
func UserLogin(db *gorm.DB) http.HandlerFunc {
	return auth.UserLogin(db)
}
func UserLogout() http.HandlerFunc {
	return auth.UserLogout()
}

func PropertyCreate(db *gorm.DB) http.HandlerFunc {
	return property.PropertyCreate(db)
}
func PropertyDelete(db *gorm.DB) http.HandlerFunc {
	return property.PropertyDelete(db)
}
func PropertyEdit(db *gorm.DB) http.HandlerFunc {
	return property.PropertyEdit(db)
}
