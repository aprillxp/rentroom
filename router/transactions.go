package router

import (
	"rentroom/handlers"
	"rentroom/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterTransactionRoutes(r *mux.Router, db *gorm.DB) {
	r.Handle("/api/transaction/create", middleware.JwtAuthUser(handlers.TransactionCreate(db))).Methods("POST")
}
