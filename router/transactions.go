package router

import (
	"rentroom/handlers"
	"rentroom/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterTransactionRoutes(r *mux.Router, db *gorm.DB) {
	r.Handle("/api/transaction/create", middleware.JwtAuthUser(handlers.TransactionCreate(db))).Methods("POST")
	r.Handle("/api/transaction/list", middleware.JwtAuthUser(handlers.TransactionList(db))).Methods("GET")
	r.Handle("/api/transaction/get/{transaction-id}", middleware.JwtAuthUser(handlers.TransactionGet(db))).Methods("GET")
}
