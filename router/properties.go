package router

import (
	"rentroom/handlers"
	"rentroom/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterPropertyRoutes(r *mux.Router, db *gorm.DB) {
	r.HandleFunc("/api/property/list", handlers.PropertyList(db)).Methods("GET")
	r.HandleFunc("/api/property/get/{property-id}", handlers.PropertyGet(db)).Methods("GET")
	r.Handle("/api/property/create", middleware.JwtAuthUser(handlers.PropertyCreate(db))).Methods("POST")
	r.Handle("/api/property/delete/{property-id}", middleware.JwtAuthUser(handlers.PropertyDelete(db))).Methods("DELETE")
	r.Handle("/api/property/edit/{property-id}", middleware.JwtAuthUser(handlers.PropertyEdit(db))).Methods("PATCH")
}
