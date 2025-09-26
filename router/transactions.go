package router

import (
	"rentroom/handlers"
	"rentroom/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterTransactionRoutes(r *mux.Router, db *gorm.DB) {
	r.Handle("/api/transaction/user/create", middleware.JwtAuthUser(handlers.TransactionCreate(db))).Methods("POST")
	r.Handle("/api/transaction/user/get/{transaction-id}", middleware.JwtAuthUser(handlers.TransactionUserGet(db))).Methods("GET")
	r.Handle("/api/transaction/user/list", middleware.JwtAuthUser(handlers.TransactionUserList(db))).Methods("GET")
	r.Handle("/api/transaction/user/cancel/{transaction-id}", middleware.JwtAuthUser(handlers.TransactionCancel(db))).Methods("PATCH")

	r.Handle("/api/transaction/tenant/list", middleware.JwtAuthUser(handlers.TransactionTenantList(db))).Methods("GET")
	r.Handle("/api/transaction/tenant/get/{transaction-id}", middleware.JwtAuthUser(handlers.TransactionTenantGet(db))).Methods("GET")
	r.Handle("/api/transaction/tenant/approve/{transaction-id}", middleware.JwtAuthUser(handlers.TransactionApprove(db))).Methods("PATCH")
	r.Handle("/api/transaction/tenant/reject/{transaction-id}", middleware.JwtAuthUser(handlers.TransactionReject(db))).Methods("PATCH")
}
