package user

import (
	"net/http"
	"rentroom/middleware"
	"rentroom/models"
	"rentroom/utils"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func TransactionUserCancel(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// AUTH
		userID, err := middleware.MustUserID(r)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}
		vars := mux.Vars(r)
		transactionID, err := strconv.ParseUint(vars["transaction-id"], 10, 64)
		if err != nil {
			utils.JSONError(w, "invalid transaction id", http.StatusBadRequest)
			return
		}
		err = utils.TransactionUserChecker(db, userID, uint(transactionID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = utils.TransactionIsPending(db, uint(transactionID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}

		// QUERY
		err = db.Model(&models.Transaction{}).
			Where("id = ?", transactionID).
			Update("status", models.StatusCanceled).Error
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		transaction, err := utils.GetUserTransaction(db, userID, uint(transactionID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "transaction canceled",
			Data:    transaction,
		}, http.StatusOK)
	}
}
