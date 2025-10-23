package admin

import (
	"net/http"
	"rentroom/middleware"
	"rentroom/models"
	"rentroom/utils"

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
		userID := r.URL.Query().Get("user-id")

		// QUERY
		var transactions []models.Transaction
		query := db.Model(&models.Transaction{})
		if userID != "" {
			query = query.Where("user_id = ?", userID)
		}
		err = query.
			Find(&transactions).Error
		if err != nil {
			utils.JSONError(w, "failed to returned user transactions", http.StatusInternalServerError)
			return
		}
		transactionsUpdated := utils.ConvertTransactionsResponse(transactions)

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "user transactions returned",
			Data:    transactionsUpdated,
		}, http.StatusOK)
	}
}
