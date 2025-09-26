package handlers

import (
	"net/http"
	property "rentroom/handlers/property"
	propertyTenant "rentroom/handlers/tenant/property"
	userAuth "rentroom/handlers/user/auth"
	userProfile "rentroom/handlers/user/profile"
	transactionUser "rentroom/handlers/user/transaction"

	"gorm.io/gorm"
)

func UserRegister(db *gorm.DB) http.HandlerFunc {
	return userAuth.UserRegister(db)
}
func UserLogin(db *gorm.DB) http.HandlerFunc {
	return userAuth.UserLogin(db)
}
func UserLogout() http.HandlerFunc {
	return userAuth.UserLogout()
}
func UserEdit(db *gorm.DB) http.HandlerFunc {
	return userProfile.UserEdit(db)
}
func UserGet(db *gorm.DB) http.HandlerFunc {
	return userProfile.UserGet(db)
}

func PropertyCreate(db *gorm.DB) http.HandlerFunc {
	return propertyTenant.PropertyCreate(db)
}
func PropertyDelete(db *gorm.DB) http.HandlerFunc {
	return propertyTenant.PropertyDelete(db)
}
func PropertyEdit(db *gorm.DB) http.HandlerFunc {
	return propertyTenant.PropertyEdit(db)
}
func PropertyList(db *gorm.DB) http.HandlerFunc {
	return property.PropertyList(db)
}
func PropertyGet(db *gorm.DB) http.HandlerFunc {
	return property.PropertyGet(db)
}

func TransactionCreate(db *gorm.DB) http.HandlerFunc {
	return transactionUser.TransactionCreate(db)
}
func TransactionCancel(db *gorm.DB) http.HandlerFunc {
	return transactionUser.TransactionCancel(db)
}
func TransactionList(db *gorm.DB) http.HandlerFunc {
	return transactionUser.TransactionList(db)
}
func TransactionGet(db *gorm.DB) http.HandlerFunc {
	return transactionUser.TransactionGet(db)
}
