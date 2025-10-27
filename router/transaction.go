package router

import (
	"rentroom/internal/handlers/transaction"
	"rentroom/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterTransactionRoutes(r *mux.Router, db *gorm.DB) {
	// ADMIN
	admin := r.PathPrefix("/api/v1/admin/transactions").Subrouter()
	admin.Use(middleware.JwtAuthAdmin)
	admin.Handle("", transaction.AdminUserList(db)).Methods("GET")
	admin.Handle("/{id}", transaction.AdminUserGet(db)).Methods("GET")
	admin.Handle("/{id}/approve", transaction.AdminApprove(db)).Methods("PATCH")
	admin.Handle("/{id}/reject", transaction.AdminReject(db)).Methods("PATCH")
	admin.Handle("/{id}/done", transaction.AdminDone(db)).Methods("PATCH")

	// TENANT
	tenant := r.PathPrefix("/api/v1/tenant/transactions").Subrouter()
	tenant.Use(middleware.JwtAuthUser)
	tenant.Handle("", transaction.TenantList(db)).Methods("GET")
	tenant.Handle("/{id}", transaction.TenantGet(db)).Methods("GET")

	// USER
	user := r.PathPrefix("/api/v1/user/transactions").Subrouter()
	user.Use(middleware.JwtAuthUser)
	user.Handle("", transaction.UserList(db)).Methods("GET")
	user.Handle("", transaction.UserCreate(db)).Methods("POST")
	user.Handle("/{id}", transaction.UserGet(db)).Methods("GET")
	user.Handle("/{id}/cancel", transaction.UserCancel(db)).Methods("PATCH")
	user.Handle("/{id}/review", transaction.UserReview(db)).Methods("POST")
}
