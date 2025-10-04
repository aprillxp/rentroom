package admin

import (
	"net/http"
	"rentroom/middleware"
	"rentroom/models"
	"rentroom/utils"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func TransactionAdminUserList(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// AUTH
		err := middleware.MustAdminID(r)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}
		vars := mux.Vars(r)
		userID, err := strconv.ParseUint(vars["user-id"], 10, 64)
		if err != nil {
			utils.JSONError(w, "invalid user id", http.StatusBadRequest)
			return
		}

		// QUERY
		var transactions []models.Transaction
		err = db.
			Where("user_id = ?", userID).
			Find(&transactions).Error
		if err != nil {
			utils.JSONError(w, "failed to returned user transactions", http.StatusInternalServerError)
			return
		}
		var transactionUpdated []models.TransactionResponse
		for _, t := range transactions {
			transactionUpdated = append(transactionUpdated, models.TransactionResponse{
				ID:         t.ID,
				PropertyID: t.PropertyID,
				Price:      t.Price,
				CheckIn:    t.CheckIn,
				CheckOut:   t.CheckOut,
				Status:     t.Status,
				VoucherID:  t.VoucherID,
			})
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "user transactions returned",
			Data:    transactionUpdated,
		}, http.StatusOK)
	}
}
