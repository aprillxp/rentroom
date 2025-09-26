package router

import (
	"rentroom/handlers"
	"rentroom/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterUserRoutes(r *mux.Router, db *gorm.DB) {
	r.HandleFunc("/api/user/auth/register", handlers.UserRegister(db)).Methods("POST")
	r.HandleFunc("/api/user/auth/login", handlers.UserLogin(db)).Methods("POST")
	r.HandleFunc("/api/user/auth/logout", handlers.UserLogout()).Methods("POST")

	r.Handle("/api/user/profile/edit", middleware.JwtAuthUser(handlers.UserEdit(db))).Methods("PATCH")
	r.Handle("/api/user/profile/get", middleware.JwtAuthUser(handlers.UserGet(db))).Methods("GET")
}
