package admin

import (
	"net/http"
	"rentroom/middleware"
	"rentroom/utils"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func TransactionAdminUserGet(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// AUTH
		err := middleware.MustAdminID(r)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}
		vars := mux.Vars(r)
		transactionID, err := strconv.ParseUint(vars["id"], 10, 64)
		if err != nil {
			utils.JSONError(w, "invalid transaction id", http.StatusBadRequest)
			return
		}

		// QUERY
		transaction, err := utils.GetTransaction(db, uint(transactionID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusNotFound)
			return
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "transaction returned",
			Data:    transaction,
		}, http.StatusOK)
	}
}
