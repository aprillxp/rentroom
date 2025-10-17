package router

import (
	"rentroom/handlers"
	"rentroom/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterTransactionRoutes(r *mux.Router, db *gorm.DB) {
	// ADMIN
	admin := r.PathPrefix("/api/v1/admin/transactions").Subrouter()
	admin.Use(middleware.JwtAuthAdmin)
	admin.Handle("", handlers.TransactionAdminUserList(db)).Methods("GET")
	admin.Handle("/{id}", handlers.TransactionAdminUserGet(db)).Methods("GET")
	admin.Handle("/{id}/approve", handlers.TransactionAdminApprove(db)).Methods("PATCH")
	admin.Handle("/{id}/reject", handlers.TransactionAdminReject(db)).Methods("PATCH")
	admin.Handle("/{id}/done", handlers.TransactionAdminDone(db)).Methods("PATCH")

	// TENANT
	tenant := r.PathPrefix("/api/v1/tenant/transactions").Subrouter()
	tenant.Use(middleware.JwtAuthUser)
	tenant.Handle("", handlers.TransactionTenantList(db)).Methods("GET")
	tenant.Handle("/{id}", handlers.TransactionTenantGet(db)).Methods("GET")

	// USER
	user := r.PathPrefix("/api/v1/user/transactions").Subrouter()
	user.Use(middleware.JwtAuthUser)
	user.Handle("", handlers.TransactionUserList(db)).Methods("GET")
	user.Handle("", handlers.TransactionUserCreate(db)).Methods("POST")
	user.Handle("/{id}", handlers.TransactionUserGet(db)).Methods("GET")
	user.Handle("/{id}/cancel", handlers.TransactionUserCancel(db)).Methods("PATCH")
	user.Handle("/{id}/review", handlers.TransactionUserReview(db)).Methods("POST")
}
