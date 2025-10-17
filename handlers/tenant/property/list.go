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
		pageStr := r.URL.Query().Get("page")   // Add query page
		limitStr := r.URL.Query().Get("limit") // Add query limit

		// Parsing page and add limit
		page, err := strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			page = 1
		}
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			limit = 10
		}
		offset := (page - 1) * limit

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

		var total int64
		query.Model(&models.Property{}).Count(&total) // Count the total before limit.

		err = query.Limit(limit).Offset(offset).Find(&properties).Error // Apply pagination limit and offset
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		totalPages := (int(total) + limit - 1) / limit // Count the total pages

		response := models.PropertiesPaginatedResponse{ // Using struct response
			Items:      properties,
			Page:       &page,
			Limit:      &limit,
			TotalItems: &total,
			TotalPages: &totalPages,
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "tenant transactions returned",
			Data:    response,
		}, http.StatusOK)
	}
}
