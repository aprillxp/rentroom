package router

import (
	"rentroom/handlers"
	"rentroom/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterAdminRoutes(r *mux.Router, db *gorm.DB) {
	r.HandleFunc("/api/admin/auth/login", handlers.AdminLogin(db)).Methods("POST")

	r.Handle("/api/admin/voucher/create", middleware.JwtAuthAdmin(handlers.AdminVoucherCreate(db))).Methods("POST")
	r.Handle("/api/admin/voucher/edit/{voucher-id}", middleware.JwtAuthAdmin(handlers.AdminVoucherEdit(db))).Methods("PATCH")
}
