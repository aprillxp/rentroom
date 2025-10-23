package tenant

import (
	"net/http"
	"rentroom/middleware"
	"rentroom/models"
	"rentroom/utils"

	"gorm.io/gorm"
)

func PropertyTenantList(db *gorm.DB) http.HandlerFunc {
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
		countryID := r.URL.Query().Get("country")

		// QUERY
		propertyIDs, err := utils.GetPropertyIDs(db, userID)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusNotFound)
			return
		}
		var properties []models.Property
		query := db
		if countryID != "" {
			query = query.Where("id IN ? AND country_id = ?", propertyIDs, countryID)
		} else {
			query = query.Where("id IN ?", propertyIDs)
		}
		err = query.Find(&properties).Error
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		propertiesUpdated := utils.ConvertPropertiesResponse(properties)

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "tenant properties returned",
			Data:    propertiesUpdated,
		}, http.StatusOK)
	}
}
