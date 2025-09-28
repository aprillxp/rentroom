package admin

import (
	"net/http"
	"rentroom/middleware"
	"rentroom/models"
	"rentroom/utils"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func AdminVoucherEdit(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// AUTH
		err := middleware.MustAdminID(r)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}
		vars := mux.Vars(r)
		voucherID, err := strconv.ParseUint(vars["voucher-id"], 10, 64)
		if err != nil {
			utils.JSONError(w, "invalid voucher id", http.StatusBadRequest)
			return
		}
		var req models.VoucherEditRequest
		err = utils.BodyChecker(r, &req)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = utils.FieldChecker(req)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = utils.VoucherUniqueness(db, *req.Name)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}

		// QUERY
		voucher, err := utils.GetVoucher(db, int(voucherID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		updates := make(map[string]interface{})
		if req.Name != nil {
			updates["name"] = *req.Name
		}
		if req.Discount != nil {
			updates["discount"] = *req.Discount
		}
		if req.Quantity != nil {
			updates["quantity"] = *req.Quantity
		}
		if req.EndPeriode != nil {
			today := time.Now()
			if req.EndPeriode.After(today) && req.EndPeriode.After(voucher.EndPeriode) {
				updates["end_periode"] = *req.EndPeriode
			}
		}
		if len(updates) > 0 {
			err = db.Model(&models.Voucher{}).
				Where("id = ?", voucherID).
				Updates(updates).Error
			if err != nil {
				utils.JSONError(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "voucher updated",
			Data:    voucher,
		}, http.StatusCreated)
	}
}
