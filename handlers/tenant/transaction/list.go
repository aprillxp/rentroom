package tenant

import (
	"net/http"
	"rentroom/middleware"
	"rentroom/models"
	"rentroom/utils"

	"gorm.io/gorm"
)

func TransactionTenantList(db *gorm.DB) http.HandlerFunc {
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
		var propertyIDs []uint
		err = db.Model(&models.UserProperties{}).
			Where("user_id = ?", userID).
			Pluck("property_id", &propertyIDs).Error
		if err != nil {
			return
		}
		var transactions []models.Transaction
		err = db.
			Where("property_id IN ?", propertyIDs).
			Find(&transactions).Error
		if err != nil {
			return
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "tenant transactions returned successfully",
			Data:    transactions,
		}, http.StatusOK)
	}
}
