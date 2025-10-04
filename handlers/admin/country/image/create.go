package admin

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

func CountryAdminImageCreate(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// AUTH
		err := middleware.MustAdminID(r)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}
		vars := mux.Vars(r)
		countryID, err := strconv.ParseUint(vars["country-id"], 10, 64)
		if err != nil {
			utils.JSONError(w, "invalid country id", http.StatusBadRequest)
			return
		}

		// QUERY
		err = r.ParseMultipartForm(10 << 20)
		if err != nil {
			utils.JSONError(w, "failed to parse form", http.StatusInternalServerError)
			return
		}
		file, header, err := r.FormFile("image")
		if err != nil {
			utils.JSONError(w, "image field is required", http.StatusInternalServerError)
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

		err = db.Model(&models.Country{}).
			Where("id = ?", countryID).
			Update("path", publicPath).Error
		if err != nil {
			_ = os.Remove(fsPath)
			utils.JSONError(w, "failed to save image record", http.StatusInternalServerError)
			return
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "country image created",
			Data:    publicPath,
		}, http.StatusOK)
	}
}
