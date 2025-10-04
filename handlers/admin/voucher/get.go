package admin

import (
	"net/http"
	"rentroom/utils"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func VoucherAdminGet(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// AUTH
		vars := mux.Vars(r)
		voucherID, err := strconv.ParseUint(vars["voucher-id"], 10, 64)
		if err != nil {
			utils.JSONError(w, "invalid voucher id", http.StatusBadRequest)
			return
		}

		// QUERY
		voucher, err := utils.GetVoucher(db, int(voucherID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "voucher returned",
			Data:    voucher,
		}, http.StatusOK)
	}
}
