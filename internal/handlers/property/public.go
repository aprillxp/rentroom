package property

import (
	"net/http"
	"rentroom/internal/models"
	service "rentroom/internal/services"
	"rentroom/utils"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func PublicList(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// QUERY
		countryID := r.URL.Query().Get("country")
		var properties []models.Property
		query := db
		if countryID != "" {
			query = query.Where("country_id = ? AND status = ?", countryID, models.StatusPublished)
		} else {
			query = query.Where("status = ?", models.StatusPublished)
		}
		err := query.Find(&properties).Error
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

func PublicGet(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// AUTH
		vars := mux.Vars(r)
		propertyID, err := strconv.ParseUint(vars["id"], 10, 64)
		if err != nil {
			utils.JSONError(w, "invalid property id", http.StatusBadRequest)
			return
		}

		// QUERY
		property, err := service.GetProperty(db, int(propertyID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if property.Status != models.StatusPublished {
			utils.JSONError(w, "property is not published", http.StatusInternalServerError)
			return
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "property returned",
			Data:    property,
		}, http.StatusOK)
	}
}

func PublicImageList(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// AUTH
		vars := mux.Vars(r)
		propertyID, err := strconv.ParseUint(vars["id"], 10, 64)
		if err != nil {
			utils.JSONError(w, "invalid property id", http.StatusBadRequest)
			return
		}
		err = utils.PropertyExist(db, uint(propertyID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// QUERY
		var images []models.Image
		err = db.Where("property_id = ?", propertyID).Find(&images).Error
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		imagesResponses := make([]models.ImageResponse, 0)
		for _, img := range images {
			imagesResponses = append(imagesResponses, models.ImageResponse{
				ID:         img.ID,
				PropertyID: img.PropertyID,
				Path:       img.Path,
			})
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "images returned from property",
			Data:    imagesResponses,
		}, http.StatusOK)
	}
}
