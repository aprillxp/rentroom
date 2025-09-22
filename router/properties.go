package router

import (
	"rentroom/handlers"
	"rentroom/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterPropertyRoutes(r *mux.Router, db *gorm.DB) {
	r.Handle("/api/property/create", middleware.JwtAuthUser(handlers.PropertyCreate(db))).Methods("POST")
}
