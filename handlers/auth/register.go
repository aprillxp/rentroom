package auth

import (
	"net/http"
	"rentroom/models"
	"rentroom/utils"
	"strings"

	"gorm.io/gorm"
)

func UserRegister(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// AUTH
		var req models.UserRegisterRequest
		err := utils.BodyChecker(r, &req)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}
		req.Username = strings.ToLower(req.Username)
		req.Email = strings.ToLower(req.Email)
		err = utils.FieldChecker(req)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = utils.PasswordValidator(req.Password)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = utils.PhoneValidator(req.Phone)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = utils.UserUniqueness(db, 0, req.Username, req.Email, req.Phone)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}

		// QUERY
		hashedPassword, err := utils.HashedPassword(req.Password)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user := models.User{
			Username:   req.Username,
			Email:      req.Email,
			Phone:      req.Phone,
			Password:   hashedPassword,
			BankID:     req.BankID,
			BankNumber: req.BankNumber,
			IsTenant:   req.IsTenant,
		}
		err = db.Create(&user).Error
		if err != nil {
			utils.JSONError(w, "failed create user", http.StatusInternalServerError)
			return
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "user registered successfully",
		}, http.StatusCreated)
	}
}
