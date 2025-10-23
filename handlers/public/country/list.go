package admin

import (
	"net/http"
	"rentroom/models"
	"rentroom/utils"

	"gorm.io/gorm"
)

func CountryList(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// QUERY
		var countries []models.Country
		err := db.
			Find(&countries).Error
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "country list returned",
			Data:    countries,
		}, http.StatusOK)
	}
}
