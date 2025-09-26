package user

import (
	"net/http"
	"rentroom/middleware"
	"rentroom/models"
	"rentroom/utils"

	"gorm.io/gorm"
)

func TransactionUserList(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// AUTH
		userID, err := middleware.MustUserID(r)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// QUERY
		var transactions []models.Transaction
		err = db.
			Where("user_id = ?", userID).
			Find(&transactions).Error
		if err != nil {
			utils.JSONError(w, "failed to returned guest transactions", http.StatusInternalServerError)
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
			Message: "user transactions returned successfully",
			Data:    transactionUpdated,
		}, http.StatusOK)
	}
}
