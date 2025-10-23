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
			utils.JSONError(w, "failed to returned user transactions", http.StatusInternalServerError)
			return
		}
		transactionsUpdated := utils.ConvertTransactionsResponse(transactions)

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "transactions returned",
			Data:    transactionsUpdated,
		}, http.StatusOK)
	}
}
