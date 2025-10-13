package handlers

import (
	"net/http"
	adminAuth "rentroom/handlers/admin/auth"
	countryAdmin "rentroom/handlers/admin/country"
	countryAdminImage "rentroom/handlers/admin/country/image"
	propertyAdmin "rentroom/handlers/admin/property"
	transactionAdmin "rentroom/handlers/admin/transaction"
	voucherAdmin "rentroom/handlers/admin/voucher"
	country "rentroom/handlers/country"
	property "rentroom/handlers/property"
	propertyImage "rentroom/handlers/property/image"
	propertyTenant "rentroom/handlers/tenant/property"
	propertyTenantImage "rentroom/handlers/tenant/property/image"
	transactionTenant "rentroom/handlers/tenant/transaction"
	userAuth "rentroom/handlers/user/auth"
	userProfile "rentroom/handlers/user/profile"
	transactionUser "rentroom/handlers/user/transaction"

	"gorm.io/gorm"
)

// USER > AUTH
func UserRegister(db *gorm.DB) http.HandlerFunc {
	return userAuth.UserRegister(db)
}
func UserLogin(db *gorm.DB) http.HandlerFunc {
	return userAuth.UserLogin(db)
}

// USER > PROFILE
func UserGet(db *gorm.DB) http.HandlerFunc {
	return userProfile.UserGet(db)
}
func UserEdit(db *gorm.DB) http.HandlerFunc {
	return userProfile.UserEdit(db)
}
func UserLogout() http.HandlerFunc {
	return userAuth.UserLogout()
}

// USER > TRANSACTION
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

// ADMIN > AUTH
func AdminLogin(db *gorm.DB) http.HandlerFunc {
	return adminAuth.AdminLogin(db)
}

// ADMIN > COUNTRY
func CountryAdminCreate(db *gorm.DB) http.HandlerFunc {
	return countryAdmin.CountryAdminCreate(db)
}

// ADMIN > COUNTRY > IMAGE
func CountryAdminImageCreate(db *gorm.DB) http.HandlerFunc {
	return countryAdminImage.CountryAdminImageCreate(db)
}
func CountryAdminImageDelete(db *gorm.DB) http.HandlerFunc {
	return countryAdminImage.CountryAdminImageDelete(db)
}

// ADMIN > PROPERTY
func PropertyAdminPublish(db *gorm.DB) http.HandlerFunc {
	return propertyAdmin.PropertyAdminPublish(db)
}
func PropertyAdminDraft(db *gorm.DB) http.HandlerFunc {
	return propertyAdmin.PropertyAdminDraft(db)
}
func PropertyAdminList(db *gorm.DB) http.HandlerFunc {
	return propertyAdmin.PropertyAdminList(db)
}
func PropertyAdminGet(db *gorm.DB) http.HandlerFunc {
	return propertyAdmin.PropertyAdminGet(db)
}

// ADMIN > TRANSACTION
func TransactionAdminApprove(db *gorm.DB) http.HandlerFunc {
	return transactionAdmin.TransactionAdminApprove(db)
}
func TransactionAdminReject(db *gorm.DB) http.HandlerFunc {
	return transactionAdmin.TransactionAdminReject(db)
}
func TransactionAdminDone(db *gorm.DB) http.HandlerFunc {
	return transactionAdmin.TransactionAdminDone(db)
}
func TransactionAdminUserList(db *gorm.DB) http.HandlerFunc {
	return transactionAdmin.TransactionAdminUserList(db)
}
func TransactionAdminUserGet(db *gorm.DB) http.HandlerFunc {
	return transactionAdmin.TransactionAdminUserGet(db)
}

// ADMIN > VOUCHER
func VoucherAdminCreate(db *gorm.DB) http.HandlerFunc {
	return voucherAdmin.VoucherAdminCreate(db)
}
func VoucherAdminEdit(db *gorm.DB) http.HandlerFunc {
	return voucherAdmin.VoucherAdminEdit(db)
}
func VoucherAdminList(db *gorm.DB) http.HandlerFunc {
	return voucherAdmin.VoucherAdminList(db)
}
func VoucherAdminGet(db *gorm.DB) http.HandlerFunc {
	return voucherAdmin.VoucherAdminGet(db)
}

// TENANT > PROPERTY
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

// TENANT > PROPERTY > IMAGE
func PropertyTenantImageList(db *gorm.DB) http.HandlerFunc {
	return propertyTenantImage.PropertyTenantImageList(db)
}
func PropertyTenantImageCreate(db *gorm.DB) http.HandlerFunc {
	return propertyTenantImage.PropertyTenantImageCreate(db)
}
func PropertyTenantImageDelete(db *gorm.DB) http.HandlerFunc {
	return propertyTenantImage.PropertyTenantImageDelete(db)
}

// TENANT > TRANSACTION
func TransactionTenantList(db *gorm.DB) http.HandlerFunc {
	return transactionTenant.TransactionTenantList(db)
}
func TransactionTenantGet(db *gorm.DB) http.HandlerFunc {
	return transactionTenant.TransactionTenantGet(db)
}

// COUNTRY
func CountryList(db *gorm.DB) http.HandlerFunc {
	return country.CountryList(db)
}
func CountryGet(db *gorm.DB) http.HandlerFunc {
	return country.CountryGet(db)
}

// PROPERTY
func PropertyList(db *gorm.DB) http.HandlerFunc {
	return property.PropertyList(db)
}
func PropertyGet(db *gorm.DB) http.HandlerFunc {
	return property.PropertyGet(db)
}

// PROPERTY > IMAGE
func PropertyImageList(db *gorm.DB) http.HandlerFunc {
	return propertyImage.PropertyImageList(db)
}
