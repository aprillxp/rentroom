package properties

import (
	"net/http"
	"rentroom/middleware"
	"rentroom/models"
	"rentroom/utils"

	"gorm.io/gorm"
)

func PropertyCreate(db *gorm.DB) http.HandlerFunc {
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
				Amenities:        req.Amenities,
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
			return nil
		})
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		propertyUpdated, err := utils.GetProperty(db, int(property.ID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "property created successfully",
			Data:    propertyUpdated,
		}, http.StatusCreated)
	}
}
