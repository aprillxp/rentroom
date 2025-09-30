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

func TransactionUserReview(db *gorm.DB) http.HandlerFunc {
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
		var req models.ReviewRequest
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
		err = utils.TransactionUserChecker(db, userID, uint(transactionID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = utils.TransactionIsDone(db, uint(transactionID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = utils.ReviewUniqueness(db, uint(transactionID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}

		// QUERY
		transaction, err := utils.GetTransaction(db, uint(transactionID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		review := models.Review{
			PropertyID:    transaction.PropertyID,
			TransactionID: uint(transactionID),
			Rating:        req.Rating,
			Description:   req.Description,
		}
		err = db.Create(&review).Error
		if err != nil {
			utils.JSONError(w, "failed create review", http.StatusInternalServerError)
			return
		}

		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "review created",
			Data:    review,
		}, http.StatusCreated)
	}
}
