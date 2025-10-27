package router

import (
	"rentroom/internal/handlers/voucher"
	"rentroom/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterVoucherRoutes(r *mux.Router, db *gorm.DB) {
	// ADMIN
	admin := r.PathPrefix("/api/v1/admin/vouchers").Subrouter()
	admin.Use(middleware.JwtAuthAdmin)
	admin.Handle("", voucher.AdminList(db)).Methods("GET")
	admin.Handle("", voucher.AdminCreate(db)).Methods("POST")
	admin.Handle("/{id}", voucher.AdminGet(db)).Methods("GET")
	admin.Handle("/{id}", voucher.AdminEdit(db)).Methods("PATCH")
}
