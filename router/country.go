package router

import (
	"rentroom/internal/handlers/country"
	"rentroom/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterCountryRoutes(r *mux.Router, db *gorm.DB) {
	// ADMIN
	admin := r.PathPrefix("/api/v1/admin/countries").Subrouter()
	admin.Use(middleware.JwtAuthAdmin)
	admin.Handle("", country.AdminList(db)).Methods("GET")
	admin.Handle("", country.AdminCreate(db)).Methods("POST")
	admin.Handle("/{id}", country.AdminGet(db)).Methods("GET")
	admin.Handle("/{id}", country.AdminDelete(db)).Methods("DELETE")
	admin.Handle("/{id}/images", country.AdminImageCreate(db)).Methods("POST")
	admin.Handle("/{id}/images", country.AdminImageDelete(db)).Methods("DELETE")
	
	// PUBLIC
	public := r.PathPrefix("/api/v1/public/countries").Subrouter()
	public.HandleFunc("", country.PublicList(db)).Methods("GET")
	public.HandleFunc("/{id}", country.PublicGet(db)).Methods("GET")
}