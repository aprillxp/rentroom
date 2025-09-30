package handlers

import (
	"net/http"
	adminAuth "rentroom/handlers/admin/auth"
	adminProperty "rentroom/handlers/admin/property"
	transactionAdmin "rentroom/handlers/admin/transaction"
	adminVoucher "rentroom/handlers/admin/voucher"
	propertyUser "rentroom/handlers/property"
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
func UserGet(db *gorm.DB) http.HandlerFunc {
	return userProfile.UserGet(db)
}
func UserEdit(db *gorm.DB) http.HandlerFunc {
	return userProfile.UserEdit(db)
}
func UserLogout() http.HandlerFunc {
	return userAuth.UserLogout()
}

// ADMIN
func AdminLogin(db *gorm.DB) http.HandlerFunc {
	return adminAuth.AdminLogin(db)
}

// PROPERTY
func PropertyAdminPublish(db *gorm.DB) http.HandlerFunc {
	return adminProperty.PropertyAdminPublish(db)
}
func PropertyAdminDraft(db *gorm.DB) http.HandlerFunc {
	return adminProperty.PropertyAdminDraft(db)
}

func PropertyTenantCreate(db *gorm.DB) http.HandlerFunc {
	return propertyTenant.PropertyTenantCreate(db)
}
func PropertyTenantDelete(db *gorm.DB) http.HandlerFunc {
	return propertyTenant.PropertyTenantDelete(db)
}
func PropertyTenantEdit(db *gorm.DB) http.HandlerFunc {
	return propertyTenant.PropertyTenantEdit(db)
}
func PropertyTenantList(db *gorm.DB) http.HandlerFunc {
	return propertyTenant.PropertyTenantList(db)
}
func PropertyTenantGet(db *gorm.DB) http.HandlerFunc {
	return propertyTenant.PropertyTenantGet(db)
}

func PropertyUserList(db *gorm.DB) http.HandlerFunc {
	return propertyUser.PropertyUserList(db)
}
func PropertyUserGet(db *gorm.DB) http.HandlerFunc {
	return propertyUser.PropertyUserGet(db)
}

// VOUCHER
func VoucherAdminCreate(db *gorm.DB) http.HandlerFunc {
	return adminVoucher.VoucherAdminCreate(db)
}
func VoucherAdminEdit(db *gorm.DB) http.HandlerFunc {
	return adminVoucher.VoucherAdminEdit(db)
}

// TRANSACTION
func TransactionUserCreate(db *gorm.DB) http.HandlerFunc {
	return transactionUser.TransactionUserCreate(db)
}
func TransactionUserCancel(db *gorm.DB) http.HandlerFunc {
	return transactionUser.TransactionUserCancel(db)
}
func TransactionUserReview(db *gorm.DB) http.HandlerFunc {
	return transactionUser.TransactionUserReview(db)
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

func TransactionAdminApprove(db *gorm.DB) http.HandlerFunc {
	return transactionAdmin.TransactionAdminApprove(db)
}
func TransactionAdminReject(db *gorm.DB) http.HandlerFunc {
	return transactionAdmin.TransactionAdminReject(db)
}
func TransactionAdminDone(db *gorm.DB) http.HandlerFunc {
	return transactionAdmin.TransactionAdminDone(db)
}
