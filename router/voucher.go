package router

import (
	"rentroom/handlers"
	"rentroom/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterVoucherRoutes(r *mux.Router, db *gorm.DB) {
	// ADMIN
	admin := r.PathPrefix("/api/v1/admin/vouchers").Subrouter()
	admin.Use(middleware.JwtAuthAdmin)
	admin.Handle("", handlers.VoucherAdminList(db)).Methods("GET")
	admin.Handle("", handlers.VoucherAdminCreate(db)).Methods("POST")
	admin.Handle("/{id}", handlers.VoucherAdminGet(db)).Methods("GET")
	admin.Handle("/{id}", handlers.VoucherAdminEdit(db)).Methods("PATCH")
}
