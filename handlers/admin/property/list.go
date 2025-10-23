package admin

import (
	"net/http"
	"rentroom/models"
	"rentroom/utils"
	"strconv"

	"gorm.io/gorm"
)

func PropertyAdminList(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// QUERY
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
		if err != nil || limit < 1 {
			limit = 10
		}
		offset := (page - 1) * limit

		var properties []models.Property
		query := db.Model(&models.Property{})
		if countryID != "" {
			query = query.Where("country_id = ?", countryID)
		}

		// ADD (Count the total before limit)
		var total int64
		query.Model(&models.Property{}).Count(&total)

		// MODIFIED (Apply pagination limit and offset)
		err = query.Offset(offset).Limit(limit).Find(&properties).Error
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		propertiesUpdated := utils.ConvertPropertiesResponse(properties)

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "properties returned",
			Data:    propertiesUpdated,
		}, http.StatusOK)
	}
}
