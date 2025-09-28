package tenant

import (
	"net/http"
	"rentroom/middleware"
	"rentroom/models"
	"rentroom/utils"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func PropertyEdit(db *gorm.DB) http.HandlerFunc {
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
		vars := mux.Vars(r)
		propertyID, err := strconv.ParseUint(vars["property-id"], 10, 64)
		if err != nil {
			utils.JSONError(w, "invalid property id", http.StatusBadRequest)
			return
		}
		err = utils.PropertyUserChecker(db, userID, uint(propertyID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}
		var req models.PropertyEditRequest
		err = utils.BodyChecker(r, &req)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = utils.FieldChecker(req)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = utils.PropertyHaveAnActiveTransaction(db, uint(propertyID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// QUERY
		updates := make(map[string]interface{})
		if req.Name != nil {
			updates["name"] = *req.Name
		}
		if req.CountryID != nil {
			updates["country_id"] = *req.CountryID
		}
		if req.Guests != nil {
			updates["guests"] = *req.Guests
		}
		if req.Price != nil {
			updates["price"] = *req.Price
		}
		updates["status"] = models.StatusDraft
		if req.DisabledDateFrom != nil {
			updates["disabled_date_from"] = *req.DisabledDateFrom
		}
		if req.DisabledDateTo != nil {
			updates["disabled_date_to"] = *req.DisabledDateTo
		}
		if req.Description != nil {
			updates["description"] = *req.Description
		}
		if req.Geo != nil {
			updates["geo"] = *req.Geo
		}
		if req.Province != nil {
			updates["province"] = *req.Province
		}
		if req.District != nil {
			updates["district"] = *req.District
		}
		if req.City != nil {
			updates["city"] = *req.City
		}
		if req.Address != nil {
			updates["address"] = *req.Address
		}
		if req.Zip != nil {
			updates["zip"] = *req.Zip
		}
		if req.Amenities != nil {
			updates["amenities"] = *req.Amenities
		}
		if len(updates) > 0 {
			err = db.Model(&models.Property{}).
				Where("id = ?", propertyID).
				Updates(updates).Error
			if err != nil {
				utils.JSONError(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		propertyUpdated, err := utils.GetProperty(db, int(propertyID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "property updated",
			Data:    propertyUpdated,
		}, http.StatusOK)
	}
}
