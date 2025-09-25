package handlers

import (
	"net/http"
	property "rentroom/handlers/properties"
	transaction "rentroom/handlers/transactions"
	auth "rentroom/handlers/users"

	"gorm.io/gorm"
)

func UserRegister(db *gorm.DB) http.HandlerFunc {
	return auth.UserRegister(db)
}
func UserEdit(db *gorm.DB) http.HandlerFunc {
	return auth.UserEdit(db)
}
func UserGet(db *gorm.DB) http.HandlerFunc {
	return auth.UserGet(db)
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
func PropertyList(db *gorm.DB) http.HandlerFunc {
	return property.PropertyList(db)
}
func PropertyGet(db *gorm.DB) http.HandlerFunc {
	return property.PropertyGet(db)
}

func TransactionCreate(db *gorm.DB) http.HandlerFunc {
	return transaction.TransactionCreate(db)
}
