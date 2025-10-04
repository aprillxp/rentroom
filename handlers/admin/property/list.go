package admin

import (
	"net/http"
	"rentroom/models"
	"rentroom/utils"

	"gorm.io/gorm"
)

func PropertyAdminList(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// QUERY
		countryID := r.URL.Query().Get("country")
		var properties []models.Property
		query := db.Model(&models.Property{})
		if countryID != "" {
			query = query.Where("country_id = ?", countryID)
		}
		err := query.Find(&properties).Error
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "properties returned",
			Data:    properties,
		}, http.StatusOK)
	}
}
