package router

import (
	"rentroom/internal/handlers/admin"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterAdminRoutes(r *mux.Router, db *gorm.DB) {
	// AUTH
	auth := r.PathPrefix("/api/v1/admin/auth").Subrouter()
	auth.HandleFunc("/login", admin.Login(db)).Methods("POST")
}