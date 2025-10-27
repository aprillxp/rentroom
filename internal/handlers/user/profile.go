package user

import (
	"net/http"
	"rentroom/internal/models"
	service "rentroom/internal/services"
	"rentroom/middleware"
	"rentroom/utils"
	"strings"

	"gorm.io/gorm"
)

func Get(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// AUTH
		userID, err := middleware.MustUserID(r)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// QUERY
		user, err := service.GetUser(db, userID)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "user returned",
			Data:    user,
		}, http.StatusOK)
	}
}

func Edit(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// AUTH
		userID, err := middleware.MustUserID(r)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}
		var req models.UserEditRequest
		err = utils.BodyChecker(r, &req)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}
		if req.Username != nil {
			*req.Username = strings.ToLower(*req.Username)
		}
		if req.Email != nil {
			*req.Email = strings.ToLower(*req.Email)
		}
		err = utils.FieldChecker(req)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}
		if req.Password != nil {
			err = utils.PasswordValidator(*req.Password)
			if err != nil {
				utils.JSONError(w, err.Error(), http.StatusBadRequest)
				return
			}
		}
		if req.Phone != nil {
			err = utils.PhoneValidator(*req.Phone)
			if err != nil {
				utils.JSONError(w, err.Error(), http.StatusBadRequest)
				return
			}
		}
		err = utils.UserUniqueness(db, uint(userID),
			utils.PtrToStrOrEmpty(req.Username),
			utils.PtrToStrOrEmpty(req.Email),
			utils.PtrToStrOrEmpty(req.Phone),
		)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}

		// QUERY
		updates := make(map[string]interface{})
		if req.Username != nil {
			updates["username"] = *req.Username
		}
		if req.Email != nil {
			updates["email"] = *req.Email
		}
		if req.Phone != nil {
			updates["phone"] = *req.Phone
		}
		if req.Password != nil {
			hashedPassword, err := utils.HashedPassword(*req.Password)
			if err != nil {
				utils.JSONError(w, err.Error(), http.StatusInternalServerError)
				return
			}
			updates["password"] = hashedPassword
		}
		if req.BankID != nil {
			updates["bank_id"] = *req.BankID
		}
		if req.BankNumber != nil {
			updates["bank_number"] = *req.BankNumber
		}
		if len(updates) > 0 {
			err = db.Model(&models.User{}).
				Where("id = ?", userID).
				Updates(updates).Error
			if err != nil {
				utils.JSONError(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		userUpdated, err := service.GetUser(db, userID)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "user updated",
			Data:    userUpdated,
		}, http.StatusOK)
	}
}
