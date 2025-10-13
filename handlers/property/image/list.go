package property

import (
	"net/http"
	"rentroom/models"
	"rentroom/utils"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func PropertyImageList(db *gorm.DB) http.HandlerFunc {
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
		var imagesResponses []models.ImageResponse
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
