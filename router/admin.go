package router

import (
	"rentroom/handlers"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterAdminRoutes(r *mux.Router, db *gorm.DB) {
	// AUTH
	auth := r.PathPrefix("/api/v1/admin/auth").Subrouter()
	auth.HandleFunc("/login", handlers.AdminLogin(db)).Methods("POST")
}
