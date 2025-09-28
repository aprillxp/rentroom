package user

import (
	"errors"
	"fmt"
	"net/http"
	"rentroom/models"
	"rentroom/utils"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func UserLogin(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// AUTH
		var req models.UserLoginRequest
		err := utils.BodyChecker(r, &req)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}
		req.Identifier = strings.ToLower(req.Identifier)
		err = utils.FieldChecker(req)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}

		// QUERY
		var user models.User
		err = db.
			Where("username = ? OR email = ? OR phone = ?", req.Identifier, req.Identifier, req.Identifier).
			First(&user).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				utils.JSONError(w, "invalid credentials", http.StatusUnauthorized)
				return
			}
			utils.JSONError(w, "database error", http.StatusInternalServerError)
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
		if err != nil {
			utils.JSONError(w, "invalid credentials", http.StatusUnauthorized)
			return
		}
		token, err := utils.GenerateJWT(uint(user.ID), "user")
		if err != nil {
			utils.JSONError(w, "token generation failed", http.StatusInternalServerError)
			return
		}
		err = utils.RedisUser.Set(utils.Ctx,
			"session:user:"+fmt.Sprint(user.ID),
			token,
			24*time.Hour,
		).Err()
		if err != nil {
			utils.JSONError(w, "redis", http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "jwt_token_user",
			Value:    token,
			HttpOnly: true,
			Secure:   false,
			Path:     "/",
			SameSite: http.SameSiteLaxMode,
		})

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "user logged in",
			Data: models.UserLoginResponse{
				Token: token,
			},
		}, http.StatusCreated)
	}
}
