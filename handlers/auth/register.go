package auth

import (
	"net/http"
	"regexp"
	"rentroom/models"
	"rentroom/utils"
	"strings"

	"gorm.io/gorm"
)

func UserRegister(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// AUTH
		var req models.User
		err := utils.BodyChecker(r, &req)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}
		req.Username = strings.ToLower(req.Username)
		req.Email = strings.ToLower(req.Email)
		if req.Username == "" || req.Password == "" || req.Email == "" || req.Phone == "" || req.BankName == 0 || req.BankNumber == 0 || req.IsTenant == 0 {
			utils.JSONError(w, "all fields required", http.StatusBadRequest)
			return
		}
		if len(req.Password) < 8 {
			utils.JSONError(w, "password must be at least 8 chars", http.StatusBadRequest)
			return
		}
		phone := utils.NormalizePhone(req.Phone)
		uppercase := regexp.MustCompile(`[A-Z]`)
		number := regexp.MustCompile(`[0-9]`)
		if !uppercase.MatchString(req.Password) {
			utils.JSONError(w, "password must contain at least one uppercase letter", http.StatusBadRequest)
			return
		}
		if !number.MatchString(req.Password) {
			utils.JSONError(w, "password must contain at least one number", http.StatusBadRequest)
			return
		}
		if !strings.Contains(req.Email, "@") {
			utils.JSONError(w, "email must contain @ symbol", http.StatusBadRequest)
			return
		}
		err = utils.UserUniqueness(db, req.Username, req.Email, req.Phone)
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
			Password:   hashedPassword,
			BankName:   req.BankName,
			BankNumber: req.BankNumber,
			IsTenant:   req.IsTenant,
			Phone:      phone,
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
