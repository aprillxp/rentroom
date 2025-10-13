package tenant

import (
	"io"
	"net/http"
	"os"
	"rentroom/middleware"
	"rentroom/models"
	"rentroom/utils"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func PropertyTenantImageCreate(db *gorm.DB) http.HandlerFunc {
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
		propertyID, err := strconv.ParseUint(vars["id"], 10, 64)
		if err != nil {
			utils.JSONError(w, "invalid property id", http.StatusBadRequest)
			return
		}
		err = utils.PropertyUserChecker(db, userID, uint(propertyID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// QUERY
		err = r.ParseMultipartForm(10 << 20)
		if err != nil {
			utils.JSONError(w, "failed to parse form", http.StatusInternalServerError)
			return
		}
		files := r.MultipartForm.File["images"]
		var uploaded []models.Image
		for _, header := range files {
			file, err := header.Open()
			if err != nil {
				utils.JSONError(w, "failed to open file", http.StatusInternalServerError)
				return
			}
			defer file.Close()

			fsPath, publicPath := utils.PathImage(header)

			out, err := os.Create(fsPath)
			if err != nil {
				utils.JSONError(w, "failed to save file", http.StatusInternalServerError)
				return
			}
			defer out.Close()

			_, err = io.Copy(out, file)
			if err != nil {
				utils.JSONError(w, "failed to write file", http.StatusInternalServerError)
				return
			}

			img := models.Image{
				PropertyID: uint(propertyID),
				Path:       publicPath,
			}
			err = db.Create(&img).Error
			if err != nil {
				_ = os.Remove(fsPath)
				utils.JSONError(w, "failed to save image record", http.StatusInternalServerError)
				return
			}

			uploaded = append(uploaded, img)
		}
		var uploadedResponses []models.ImageResponse
		for _, img := range uploaded {
			uploadedResponses = append(uploadedResponses, models.ImageResponse{
				ID:         img.ID,
				PropertyID: img.PropertyID,
				Path:       img.Path,
			})
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "images added to property",
			Data:    uploadedResponses,
		}, http.StatusCreated)
	}
}
