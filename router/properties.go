package router

import (
	"rentroom/handlers"
	"rentroom/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterPropertyRoutes(r *mux.Router, db *gorm.DB) {
	// PUBLIC
	public := r.PathPrefix("/api/v1/public/properties").Subrouter()
	public.HandleFunc("", handlers.PropertyList(db)).Methods("GET")
	public.HandleFunc("/{id}", handlers.PropertyGet(db)).Methods("GET")
	public.HandleFunc("/{id}/images", handlers.PropertyImageList(db)).Methods("GET")

	// TENANT
	tenant := r.PathPrefix("/api/v1/tenant/properties").Subrouter()
	tenant.Use(middleware.JwtAuthUser)
	tenant.Handle("", handlers.PropertyTenantList(db)).Methods("GET")
	tenant.Handle("", handlers.PropertyTenantCreate(db)).Methods("POST")
	tenant.Handle("/{id}", handlers.PropertyTenantGet(db)).Methods("GET")
	tenant.Handle("/{id}", handlers.PropertyTenantEdit(db)).Methods("PATCH")
	tenant.Handle("/{id}", handlers.PropertyTenantDelete(db)).Methods("DELETE")
	tenant.Handle("/{id}/images", handlers.PropertyTenantImageList(db)).Methods("GET")
	tenant.Handle("/{id}/images", handlers.PropertyTenantImageCreate(db)).Methods("POST")
	tenant.Handle("/{id}/images", handlers.PropertyTenantImageDelete(db)).Methods("DELETE")

	// ADMIN
	admin := r.PathPrefix("/api/v1/admin/properties").Subrouter()
	admin.Use(middleware.JwtAuthAdmin)
	admin.Handle("", handlers.PropertyAdminList(db)).Methods("GET")
	admin.Handle("/{id}", handlers.PropertyAdminGet(db)).Methods("GET")
	admin.Handle("/{id}/publish", handlers.PropertyAdminPublish(db)).Methods("PATCH")
	admin.Handle("/{id}/draft", handlers.PropertyAdminDraft(db)).Methods("PATCH")
}
