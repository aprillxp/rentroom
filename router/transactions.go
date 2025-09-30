package router

import (
	"rentroom/handlers"
	"rentroom/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterTransactionRoutes(r *mux.Router, db *gorm.DB) {
	r.Handle("/api/transaction/user/create", middleware.JwtAuthUser(handlers.TransactionUserCreate(db))).Methods("POST")
	r.Handle("/api/transaction/user/cancel/{transaction-id}", middleware.JwtAuthUser(handlers.TransactionUserCancel(db))).Methods("PATCH")
	r.Handle("/api/transaction/user/review/{transaction-id}", middleware.JwtAuthUser(handlers.TransactionUserReview(db))).Methods("POST")
	r.Handle("/api/transaction/user/get/{transaction-id}", middleware.JwtAuthUser(handlers.TransactionUserGet(db))).Methods("GET")
	r.Handle("/api/transaction/user/list", middleware.JwtAuthUser(handlers.TransactionUserList(db))).Methods("GET")

	r.Handle("/api/transaction/tenant/list", middleware.JwtAuthUser(handlers.TransactionTenantList(db))).Methods("GET")
	r.Handle("/api/transaction/tenant/get/{transaction-id}", middleware.JwtAuthUser(handlers.TransactionTenantGet(db))).Methods("GET")

	r.Handle("/api/transaction/admin/approve/{transaction-id}", middleware.JwtAuthAdmin(handlers.TransactionAdminApprove(db))).Methods("PATCH")
	r.Handle("/api/transaction/admin/reject/{transaction-id}", middleware.JwtAuthAdmin(handlers.TransactionAdminReject(db))).Methods("PATCH")
	r.Handle("/api/transaction/admin/done/{transaction-id}", middleware.JwtAuthAdmin(handlers.TransactionAdminDone(db))).Methods("PATCH")
}
