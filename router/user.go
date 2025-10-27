package router

import (
	"rentroom/internal/handlers/user"
	"rentroom/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterUserRoutes(r *mux.Router, db *gorm.DB) {
	// AUTH
	auth := r.PathPrefix("/api/v1/user/auth").Subrouter()
	auth.HandleFunc("/register", user.Register(db)).Methods("POST")
	auth.HandleFunc("/login", user.Login(db)).Methods("POST")
	auth.HandleFunc("/logout", user.Logout()).Methods("POST")

	// PROFILE
	profile := r.PathPrefix("/api/v1/user/profile").Subrouter()
	profile.Use(middleware.JwtAuthUser)
	profile.Handle("", user.Get(db)).Methods("GET")
	profile.Handle("", user.Edit(db)).Methods("PATCH")	
}
