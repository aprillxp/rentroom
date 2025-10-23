package router

import (
	"rentroom/handlers"
	"rentroom/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterCountryRoutes(r *mux.Router, db *gorm.DB) {
	// ADMIN
	admin := r.PathPrefix("/api/v1/admin/countries").Subrouter()
	admin.Use(middleware.JwtAuthAdmin)
	admin.Handle("", handlers.CountryAdminList(db)).Methods("GET")
	admin.Handle("", handlers.CountryAdminCreate(db)).Methods("POST")
	admin.Handle("/{id}", handlers.CountryAdminGet(db)).Methods("GET")
	admin.Handle("/{id}", handlers.CountryAdminDelete(db)).Methods("DELETE")
	admin.Handle("/{id}/images", handlers.CountryAdminImageCreate(db)).Methods("POST")
	admin.Handle("/{id}/images", handlers.CountryAdminImageDelete(db)).Methods("DELETE")
	
	// PUBLIC
	public := r.PathPrefix("/api/v1/public/countries").Subrouter()
	public.HandleFunc("", handlers.CountryList(db)).Methods("GET")
	public.HandleFunc("/{id}", handlers.CountryGet(db)).Methods("GET")
}
