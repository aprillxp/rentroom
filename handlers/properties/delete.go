package properties

import (
	"net/http"
	"rentroom/middleware"
	"rentroom/models"
	"rentroom/utils"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func PropertyDelete(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// AUTH
		userID, err := middleware.MustUserID(r)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}
		vars := mux.Vars(r)
		propertyID, err := strconv.ParseUint(vars["property-id"], 10, 64)
		if err != nil {
			utils.JSONError(w, "invalid property id", http.StatusBadRequest)
			return
		}
		err = utils.PropertyUserChecker(db, userID, int(propertyID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// QUERY
		err = db.Transaction(func(tx *gorm.DB) error {
			err = tx.Delete(&models.Property{}, propertyID).Error
			if err != nil {
				return err
			}
			err = tx.Where("property_id = ?", propertyID).Delete(&models.UserProperties{}).Error
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "property deleted successfully",
		}, http.StatusOK)
	}
}
