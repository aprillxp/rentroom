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
	// r.HandleFunc("/api/admin/general/voucher/edit", handlers.(db)).Methods("PATCH")
	// r.HandleFunc("/api/admin/general/voucher/delete", handlers.(db)).Methods("DELETE")
}
