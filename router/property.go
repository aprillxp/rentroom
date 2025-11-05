package router

import (
	"rentroom/internal/handlers/property"
	"rentroom/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterPropertyRoutes(r *mux.Router, db *gorm.DB) {
	// ADMIN
	admin := r.PathPrefix("/api/v1/admin/properties").Subrouter()
	admin.Use(middleware.JwtAuthAdmin)
	admin.Handle("", property.AdminList(db)).Methods("GET")
	admin.Handle("/{id}", property.AdminGet(db)).Methods("GET")
	admin.Handle("/{id}/publish", property.AdminPublish(db)).Methods("PATCH")
	admin.Handle("/{id}/draft", property.AdminDraft(db)).Methods("PATCH")

	// TENANT
	tenant := r.PathPrefix("/api/v1/tenant/properties").Subrouter()
	tenant.Use(middleware.JwtAuthUser)
	tenant.Handle("", property.TenantList(db)).Methods("GET")
	tenant.Handle("", property.TenantCreate(db)).Methods("POST")
	tenant.Handle("/{id}", property.TenantGet(db)).Methods("GET")
	tenant.Handle("/{id}", property.TenantEdit(db)).Methods("PATCH")
	tenant.Handle("/{id}", property.TenantDelete(db)).Methods("DELETE")
	tenant.Handle("/{id}/images", property.TenantImageList(db)).Methods("GET")
	tenant.Handle("/{id}/images", property.TenantImageCreate(db)).Methods("POST")
	tenant.Handle("/{id}/images", property.TenantImageDelete(db)).Methods("DELETE")

	// PUBLIC
	public := r.PathPrefix("/api/v1/public/properties").Subrouter()
	public.HandleFunc("", property.PublicList(db)).Methods("GET")
	public.HandleFunc("/{id}", property.PublicGet(db)).Methods("GET")
	// public.HandleFunc("/{id}/images", property.PublicImageList(db)).Methods("GET")
}
