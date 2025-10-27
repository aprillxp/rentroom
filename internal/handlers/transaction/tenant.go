package transaction

import (
	"net/http"
	"rentroom/internal/models"
	service "rentroom/internal/services"
	"rentroom/middleware"
	"rentroom/utils"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func TenantList(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// AUTH
		userID, err := middleware.MustUserID(r)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}
		err = utils.UserIsTenant(db, userID)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// QUERY
		propertyIDs, err := service.GetPropertyIDs(db, userID)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusNotFound)
			return
		}
		var transactions []models.Transaction
		err = db.
			Where("property_id IN ?", propertyIDs).
			Find(&transactions).Error
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		transactionsUpdated := utils.ConvertTransactionsResponse(transactions)

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "tenant transactions returned",
			Data:    transactionsUpdated,
		}, http.StatusOK)
	}
}

func TenantGet(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// AUTH
		userID, err := middleware.MustUserID(r)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}
		err = utils.UserIsTenant(db, userID)
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
		propertyIDs, err := service.GetPropertyIDs(db, userID)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusNotFound)
			return
		}
		var transaction models.Transaction
		err = db.
			Where("id = ? AND property_id IN ?", transactionID, propertyIDs).
			First(&transaction).Error
		if err != nil {
			utils.JSONError(w, "transaction not found", http.StatusNotFound)
			return
		}
		transactionUpdated := utils.ConvertTransactionResponse(transaction)

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "transaction returned",
			Data:    transactionUpdated,
		}, http.StatusOK)

	}
}
