package router

import (
	"rentroom/handlers"
	"rentroom/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterCountryRoutes(r *mux.Router, db *gorm.DB) {
	r.Handle("/api/country/admin/create", middleware.JwtAuthAdmin(handlers.CountryAdminCreate(db))).Methods("POST")

	r.Handle("/api/country/admin/image/create/{country-id}", middleware.JwtAuthAdmin(handlers.CountryAdminImageCreate(db))).Methods("POST")
	r.Handle("/api/country/admin/image/delete/{country-id}", middleware.JwtAuthAdmin(handlers.CountryAdminImageDelete(db))).Methods("DELETE")

	r.HandleFunc("/api/country/list", handlers.CountryList(db)).Methods("GET")
	r.HandleFunc("/api/country/get/{country-id}", handlers.CountryGet(db)).Methods("GET")
}
