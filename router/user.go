package router

import (
	"rentroom/handlers"
	"rentroom/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterUserRoutes(r *mux.Router, db *gorm.DB) {
	//AUTH
	auth := r.PathPrefix("/api/v1/user/auth").Subrouter()
	auth.HandleFunc("/register", handlers.UserRegister(db)).Methods("POST")
	auth.HandleFunc("/login", handlers.UserLogin(db)).Methods("POST")
	auth.HandleFunc("/logout", handlers.UserLogout()).Methods("POST")

	//PROFILE
	profile := r.PathPrefix("/api/v1/user/profile").Subrouter()
	profile.Use(middleware.JwtAuthUser)
	profile.Handle("", handlers.UserGet(db)).Methods("GET")
	profile.Handle("", handlers.UserEdit(db)).Methods("PATCH")
}
