package router

import (
	"rentroom/handlers"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func UserRoutes(r *mux.Router, db *gorm.DB) {
	r.HandleFunc("/api/register", handlers.UserRegister(db)).Methods("POST")
	r.HandleFunc("/api/login", handlers.UserLogin(db)).Methods("POST")
}
