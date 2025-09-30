package router

import (
	"rentroom/handlers"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterAdminRoutes(r *mux.Router, db *gorm.DB) {
	r.HandleFunc("/api/admin/auth/login", handlers.AdminLogin(db)).Methods("POST")
}
