package admin

import (
	"net/http"
	"rentroom/models"
	"rentroom/utils"
	"strconv"

	"gorm.io/gorm"
)

func CountryList(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ADD (Query params for pagination)
		pageStr := r.URL.Query().Get("page")
		limitStr := r.URL.Query().Get("limit")

		// PARSING PAGE & LIMIT
		page, err := strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			page = 1
		}
		limit, err := strconv.Atoi(limitStr)
		if err != nil || limit < 1 {
			limit = 10
		}
		offset := (page - 1) * limit

		// QUERY
		var countries []models.Country

		var total int64
		db.Model(&models.Country{}).Count(&total)

		err = db.Offset(offset).Limit(limit).Find(&countries).Error
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
