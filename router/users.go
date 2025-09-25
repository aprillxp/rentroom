package router

import (
	"rentroom/handlers"
	"rentroom/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterUserRoutes(r *mux.Router, db *gorm.DB) {
	r.HandleFunc("/api/user/register", handlers.UserRegister(db)).Methods("POST")
	r.HandleFunc("/api/user/login", handlers.UserLogin(db)).Methods("POST")
	r.HandleFunc("/api/user/logout", handlers.UserLogout()).Methods("POST")
	r.Handle("/api/user/edit", middleware.JwtAuthUser(handlers.UserEdit(db))).Methods("PATCH")
	r.Handle("/api/user/get", middleware.JwtAuthUser(handlers.UserGet(db))).Methods("GET")
}
