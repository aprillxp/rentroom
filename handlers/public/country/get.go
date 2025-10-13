package admin

import (
	"net/http"
	"rentroom/utils"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func CountryGet(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// AUTH
		vars := mux.Vars(r)
		countryID, err := strconv.ParseUint(vars["country-id"], 10, 64)
		if err != nil {
			utils.JSONError(w, "invalid country id", http.StatusBadRequest)
			return
		}

		// QUERY
		country, err := utils.GetCountry(db, int(countryID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "country returned",
			Data:    country,
		}, http.StatusOK)
	}
}
