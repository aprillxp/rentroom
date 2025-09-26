package transactions

import (
	"math"
	"net/http"
	"rentroom/middleware"
	"rentroom/models"
	"rentroom/utils"

	"gorm.io/gorm"
)

func TransactionCreate(db *gorm.DB) http.HandlerFunc {
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
		property, err := utils.GetProperty(db, int(req.PropertyID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		discount := 0.0
		var voucherID uint
		if req.VoucherID != nil {
			discount = utils.GetVoucher(db, int(*req.VoucherID))
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
		transactionUpdated, err := utils.GetTransaction(db, userID, transaction.ID)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "transaction created successfully",
			Data:    transactionUpdated,
		}, http.StatusCreated)
	}
}
