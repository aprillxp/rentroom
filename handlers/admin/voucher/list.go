package admin

import (
	"net/http"
	"rentroom/models"
	"rentroom/utils"

	"gorm.io/gorm"
)

func VoucherAdminList(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// QUERY
		var vouchers []models.Voucher
		err := db.
			Find(&vouchers).Error
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "voucher list returned",
			Data:    vouchers,
		}, http.StatusOK)
	}
}
