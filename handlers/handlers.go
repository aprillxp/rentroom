package handlers

import (
	"net/http"
	adminAuth "rentroom/handlers/admin/auth"
	adminProperty "rentroom/handlers/admin/property"
	transactionAdmin "rentroom/handlers/admin/transaction"
	adminVoucher "rentroom/handlers/admin/voucher"
	property "rentroom/handlers/property"
	propertyTenant "rentroom/handlers/tenant/property"
	transactionTenant "rentroom/handlers/tenant/transaction"
	userAuth "rentroom/handlers/user/auth"
	userProfile "rentroom/handlers/user/profile"
	transactionUser "rentroom/handlers/user/transaction"

	"gorm.io/gorm"
)

// USER
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

// ADMIN
func AdminLogin(db *gorm.DB) http.HandlerFunc {
	return adminAuth.AdminLogin(db)
}
func PropertyPublish(db *gorm.DB) http.HandlerFunc {
	return adminProperty.PropertyPublish(db)
}
func PropertyDraft(db *gorm.DB) http.HandlerFunc {
	return adminProperty.PropertyDraft(db)
}
func AdminVoucherCreate(db *gorm.DB) http.HandlerFunc {
	return adminVoucher.AdminVoucherCreate(db)
}
func AdminVoucherEdit(db *gorm.DB) http.HandlerFunc {
	return adminVoucher.AdminVoucherEdit(db)
}

// PROPERTY
func PropertyTenantList(db *gorm.DB) http.HandlerFunc {
	return propertyTenant.PropertyTenantList(db)
}
func PropertyTenantGet(db *gorm.DB) http.HandlerFunc {
	return propertyTenant.PropertyTenantGet(db)
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

// TRANSACTION
func TransactionCreate(db *gorm.DB) http.HandlerFunc {
	return transactionUser.TransactionCreate(db)
}
func TransactionCancel(db *gorm.DB) http.HandlerFunc {
	return transactionUser.TransactionCancel(db)
}
func TransactionUserList(db *gorm.DB) http.HandlerFunc {
	return transactionUser.TransactionUserList(db)
}
func TransactionUserGet(db *gorm.DB) http.HandlerFunc {
	return transactionUser.TransactionUserGet(db)
}
func TransactionTenantList(db *gorm.DB) http.HandlerFunc {
	return transactionTenant.TransactionTenantList(db)
}
func TransactionTenantGet(db *gorm.DB) http.HandlerFunc {
	return transactionTenant.TransactionTenantGet(db)
}
func TransactionApprove(db *gorm.DB) http.HandlerFunc {
	return transactionAdmin.TransactionApprove(db)
}
func TransactionReject(db *gorm.DB) http.HandlerFunc {
	return transactionAdmin.TransactionReject(db)
}
func TransactionDone(db *gorm.DB) http.HandlerFunc {
	return transactionAdmin.TransactionDone(db)
}
