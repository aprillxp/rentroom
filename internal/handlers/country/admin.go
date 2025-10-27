package country

import (
	"io"
	"net/http"
	"os"
	"rentroom/internal/models"
	service "rentroom/internal/services"
	"rentroom/middleware"
	"rentroom/utils"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func AdminList(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// AUTH
		err := middleware.MustAdminID(r)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// QUERY
		var countries []models.Country
		err = db.
			Find(&countries).Error
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

func AdminGet(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// AUTH
		err := middleware.MustAdminID(r)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}
		vars := mux.Vars(r)
		countryID, err := strconv.ParseUint(vars["id"], 10, 64)
		if err != nil {
			utils.JSONError(w, "invalid country id", http.StatusBadRequest)
			return
		}

		// QUERY
		country, err := service.GetCountry(db, int(countryID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "country returned",
			Data:    country,
		}, http.StatusOK)
	}
}

func AdminCreate(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// AUTH
		err := middleware.MustAdminID(r)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}
		var req models.CountryRequest
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
		err = utils.CountryUniqueness(db, req.Name)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}

		// QUERY
		country := models.Country{
			Name:        req.Name,
			Description: req.Description,
		}
		err = db.Create(&country).Error
		if err != nil {
			utils.JSONError(w, "failed create country", http.StatusInternalServerError)
			return
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "country created",
			Data:    country,
		}, http.StatusCreated)
	}
}

func AdminDelete(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		// AUTH
		err := middleware.MustAdminID(r)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}
		vars := mux.Vars(r)
		countryID, err := strconv.ParseUint(vars["id"], 10, 64)
		if err != nil {
			utils.JSONError(w, "invalid country id", http.StatusBadRequest)
			return
		}
		err = utils.CountryValidator(db, uint(countryID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}
		err = utils.CountryHaveProperty(db, uint(countryID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// QUERY
		err = db.Delete(&models.Country{}, countryID).Error
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "country deleted",
		}, http.StatusOK)
	}
}

func AdminImageCreate(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// AUTH
		err := middleware.MustAdminID(r)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}
		vars := mux.Vars(r)
		countryID, err := strconv.ParseUint(vars["id"], 10, 64)
		if err != nil {
			utils.JSONError(w, "invalid country id", http.StatusBadRequest)
			return
		}
		err = utils.CountryValidator(db, uint(countryID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
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

func AdminImageDelete(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// AUTH
		err := middleware.MustAdminID(r)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}
		vars := mux.Vars(r)
		countryID, err := strconv.ParseUint(vars["id"], 10, 64)
		if err != nil {
			utils.JSONError(w, "invalid country id", http.StatusBadRequest)
			return
		}
		err = utils.CountryValidator(db, uint(countryID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// QUERY
		country, err := service.GetCountry(db, int(countryID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if country.Path == nil {
			utils.JSONError(w, "country does not have image", http.StatusInternalServerError)
			return
		}
		err = os.Remove("." + *country.Path)
		if err != nil && !os.IsNotExist(err) {
			utils.JSONError(w, "failed to delete image file", http.StatusInternalServerError)
			return
		}
		err = db.
			Model(&country).
			Update("path", nil).Error
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "image deleted from country",
			Data:    country,
		}, http.StatusOK)
	}
}

