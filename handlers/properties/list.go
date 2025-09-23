package properties

import (
	"net/http"
	"rentroom/models"
	"rentroom/utils"

	"gorm.io/gorm"
)

func PropertyList(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// QUERY
		country := r.URL.Query().Get("country")
		var properties []models.Property
		query := db
		if country != "" {
			query = query.Where("country_id = ?", country)
		}
		err := query.Find(&properties).Error
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "properties returned successfully",
			Data:    properties,
		}, http.StatusOK)
	}
}
