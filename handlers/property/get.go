package property

import (
	"net/http"
	"rentroom/models"
	"rentroom/utils"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func PropertyGet(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// AUTH
		vars := mux.Vars(r)
		propertyID, err := strconv.ParseUint(vars["id"], 10, 64)
		if err != nil {
			utils.JSONError(w, "invalid property id", http.StatusBadRequest)
			return
		}

		// QUERY
		property, err := utils.GetProperty(db, int(propertyID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if property.Status != models.StatusPublished {
			utils.JSONError(w, "property is not published", http.StatusInternalServerError)
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
