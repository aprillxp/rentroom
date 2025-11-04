package user

import (
	"errors"
	"fmt"
	"net/http"
	"rentroom/internal/models"
	"rentroom/middleware"
	"rentroom/utils"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(db *gorm.DB) http.HandlerFunc {
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
			Message: "user registered",
			Data:    user,
		}, http.StatusCreated)
	}
}

func Login(db *gorm.DB) http.HandlerFunc {
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

func Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt_token_user")
		if err != nil {
			utils.JSONError(w, "no cookie", http.StatusUnauthorized)
			return
		}
		claims, err := middleware.Validate(cookie, "user")
		if err != nil {
			utils.JSONError(w, "invalid token", http.StatusUnauthorized)
			return
		}
		userID := int(claims["id"].(float64))
		redisKey := fmt.Sprintf("session:client:%d", userID)
		utils.RedisUser.Del(utils.Ctx, redisKey)

		http.SetCookie(w, &http.Cookie{
			Name:     "jwt_token_user",
			Value:    "",
			Path:     "/",
			MaxAge:   -1,
			HttpOnly: true,
			Secure:   false,
			SameSite: http.SameSiteLaxMode,
		})

		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "logged out",
		}, http.StatusOK)
	}
}
