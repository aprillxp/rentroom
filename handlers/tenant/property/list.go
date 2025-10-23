package tenant

import (
	"net/http"
	"rentroom/middleware"
	"rentroom/models"
	"rentroom/utils"
	"strconv"

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

		// ADD (Query params for pagination)
		pageStr := r.URL.Query().Get("page")
		limitStr := r.URL.Query().Get("limit")

		// ADD (Parsing page and limit)
		page, err := strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			page = 1
		}
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			limit = 10
		}
		offset := (page - 1) * limit

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

		// ADD (Count the total before limit)
		var total int64
		query.Model(&models.Property{}).Count(&total)

		// MODIFIED (Apply pagination limit and offset)
		err = query.Limit(limit).Offset(offset).Find(&properties).Error
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
