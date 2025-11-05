package router

import (
	"rentroom/internal/handlers/user"
	repository "rentroom/internal/repositories/user"
	service "rentroom/internal/services/user"
	"rentroom/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterUserRoutes(r *mux.Router, db *gorm.DB) {
	userRepo := repository.NewGormUserRepository(db)
	userService := service.NewUserService(userRepo)

	authHandler := user.NewAuthHandler(userService)
	profileHandler := user.NewProfileHandler(userService)

	// AUTH
	auth := r.PathPrefix("/api/v1/user/auth").Subrouter()
	auth.HandleFunc("/register", authHandler.Register()).Methods("POST")
	auth.HandleFunc("/login", authHandler.Login()).Methods("POST")
	auth.HandleFunc("/logout", authHandler.Logout()).Methods("POST")

	// PROFILE
	profile := r.PathPrefix("/api/v1/user/profile").Subrouter()
	profile.Use(middleware.JwtAuthUser)
	profile.Handle("", profileHandler.Get()).Methods("GET")
	profile.Handle("", profileHandler.Edit()).Methods("PATCH")
}
