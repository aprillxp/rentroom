package transaction

import (
	"math"
	"net/http"
	"rentroom/internal/models"
	service "rentroom/internal/services"
	"rentroom/middleware"
	"rentroom/utils"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func UserList(db *gorm.DB) http.HandlerFunc {
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

func UserGet(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// AUTH
		userID, err := middleware.MustUserID(r)
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
		err = utils.TransactionUserChecker(db, userID, uint(transactionID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusNotFound)
			return
		}
		transaction, err := service.GetTransaction(db, uint(transactionID))
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

func UserCreate(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// AUTH
		userID, err := middleware.MustUserID(r)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}
		var req models.TransactionRequest
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
		err = utils.PropertOwnedByUser(db, userID, uint(req.PropertyID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = utils.TransactionOwnedByUser(db, userID, uint(req.PropertyID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = utils.PropertyAvailable(db, req.PropertyID, req.CheckIn, req.CheckOut)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}

		// QUERY
		nights := int(math.Ceil(req.CheckOut.Sub(req.CheckIn).Hours() / 24))
		property, err := service.GetProperty(db, int(req.PropertyID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		discount := 0.0
		var voucherID uint
		if req.VoucherID != nil {
			voucher, err := service.GetVoucher(db, int(*req.VoucherID))
			if err != nil {
				utils.JSONError(w, err.Error(), http.StatusInternalServerError)
				return
			}
			discount = voucher.Discount
			voucherID = *req.VoucherID
		} else {
			voucherID = 0
		}
		price := (property.Price * float64(nights)) * (1.0 - discount)
		transaction := models.Transaction{
			PropertyID: req.PropertyID,
			UserID:     userID,
			Price:      price,
			CheckIn:    req.CheckIn,
			CheckOut:   req.CheckOut,
			Status:     models.StatusPending,
			VoucherID:  voucherID,
		}
		err = db.Create(&transaction).Error
		if err != nil {
			utils.JSONError(w, "failed create transaction", http.StatusInternalServerError)
			return
		}
		transactionUpdated, err := service.GetTransaction(db, transaction.ID)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusNotFound)
			return
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "transaction created",
			Data:    transactionUpdated,
		}, http.StatusCreated)
	}
}

func UserCancel(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// AUTH
		userID, err := middleware.MustUserID(r)
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
		transaction, err := service.GetTransaction(db, uint(transactionID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusNotFound)
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

func UserReview(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// AUTH
		userID, err := middleware.MustUserID(r)
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
		transaction, err := service.GetTransaction(db, uint(transactionID))
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
