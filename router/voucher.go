package router

import (
	"rentroom/handlers"
	"rentroom/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterVoucherRoutes(r *mux.Router, db *gorm.DB) {
	r.Handle("/api/voucher/admin/create", middleware.JwtAuthAdmin(handlers.VoucherAdminCreate(db))).Methods("POST")
	r.Handle("/api/voucher/admin/edit/{voucher-id}", middleware.JwtAuthAdmin(handlers.VoucherAdminEdit(db))).Methods("PATCH")
}
