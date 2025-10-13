package tenant

import (
	"net/http"
	"rentroom/middleware"
	"rentroom/utils"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func PropertyTenantGet(db *gorm.DB) http.HandlerFunc {
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
		propertyID, err := strconv.ParseUint(vars["id"], 10, 64)
		if err != nil {
			utils.JSONError(w, "invalid property id", http.StatusBadRequest)
			return
		}
		err = utils.PropertyUserChecker(db, userID, uint(propertyID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// QUERY
		property, err := utils.GetProperty(db, int(propertyID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "property returned",
			Data:    property,
		}, http.StatusOK)
	}
}
