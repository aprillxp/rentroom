package admin

import (
	"net/http"
	"os"
	"rentroom/middleware"
	"rentroom/utils"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func CountryAdminImageDelete(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// AUTH
		err := middleware.MustAdminID(r)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}
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
		if country.Path == nil {
			utils.JSONError(w, "country does not have image", http.StatusInternalServerError)
			return
		}
		err = os.Remove("." + *country.Path)
		if err != nil && !os.IsNotExist(err) {
			utils.JSONError(w, "failed to delete image file", http.StatusInternalServerError)
			return
		}
		err = db.
			Model(&country).
			Update("path", "").Error
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "iamge deleted from country",
			Data:    country,
		}, http.StatusOK)
	}
}
