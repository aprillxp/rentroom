package property

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

func TenantList(db *gorm.DB) http.HandlerFunc {
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

		// QUERY
		propertyIDs, err := service.GetPropertyIDs(db, userID)
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
		err = query.Find(&properties).Error
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

func TenantGet(db *gorm.DB) http.HandlerFunc {
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
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}

		// QUERY
		property, err := service.GetProperty(db, int(propertyID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
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

func TenantCreate(db *gorm.DB) http.HandlerFunc {
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
		var req models.PropertyCreateRequest
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
		err = utils.CountryValidator(db, uint(req.CountryID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}

		// QUERY
		var property models.Property
		err = db.Transaction(func(tx *gorm.DB) error {
			property = models.Property{
				Name:             req.Name,
				CountryID:        req.CountryID,
				Guests:           req.Guests,
				Price:            req.Price,
				Status:           models.StatusDraft,
				DisabledDateFrom: req.DisabledDateFrom,
				DisabledDateTo:   req.DisabledDateTo,
				Description:      req.Description,
				Geo:              req.Geo,
				Province:         req.Province,
				District:         req.District,
				City:             req.City,
				Address:          req.Address,
				Zip:              req.Zip,
			}
			err = tx.Create(&property).Error
			if err != nil {
				return err
			}
			userProperty := models.UserProperties{
				UserID:     uint(userID),
				PropertyID: property.ID,
			}
			err = tx.Create(&userProperty).Error
			if err != nil {
				return err
			}
			for _, amenityID := range req.Amenities {
				propertyAmenities := models.PropertyAmenities{
					PropertyID: property.ID,
					AmenityID:  amenityID,
				}
				err := tx.Create(&propertyAmenities).Error
				if err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		propertyUpdated, err := service.GetProperty(db, int(property.ID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "property created",
			Data:    propertyUpdated,
		}, http.StatusCreated)
	}
}

func TenantEdit(db *gorm.DB) http.HandlerFunc {
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
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
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
		propertyUpdated, err := service.GetProperty(db, int(propertyID))
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

func TenantDelete(db *gorm.DB) http.HandlerFunc {
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
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}
		err = utils.PropertyHaveAnActiveTransaction(db, uint(propertyID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// QUERY
		var images []models.Image
		err = db.
			Where("property_id = ?", propertyID).
			Find(&images).Error
		if err != nil {
			utils.JSONError(w, "failed to fetch images", http.StatusInternalServerError)
			return
		}
		for _, img := range images {
			_ = os.Remove("." + img.Path)
		}
		err = db.Delete(&models.Property{}, propertyID).Error
		if err != nil {
			utils.JSONError(w, "failed to delete property", http.StatusInternalServerError)
			return
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "property deleted",
		}, http.StatusOK)
	}
}

func TenantImageList(db *gorm.DB) http.HandlerFunc {
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
		err = utils.PropertyUserChecker(db, uint(userID), uint(propertyID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
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

func TenantImageCreate(db *gorm.DB) http.HandlerFunc {
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
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
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

func TenantImageDelete(db *gorm.DB) http.HandlerFunc {
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
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}
		var req models.ImageDeleteRequest
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

		// QUERY
		var images []models.Image
		err = db.
			Where("property_id = ? AND id IN ?", propertyID, req.ImageIds).
			First(&images).Error
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for _, img := range images {
			_ = os.Remove("." + img.Path)
		}
		err = db.
			Where("property_id = ? AND id IN ?", propertyID, req.ImageIds).
			Delete(&models.Image{}).Error
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "images deleted from property",
		}, http.StatusCreated)
	}
}