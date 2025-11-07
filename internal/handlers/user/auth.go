package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rentroom/internal/models"
	services "rentroom/internal/services/user"
	"rentroom/middleware"
	"rentroom/utils"
	"strings"
	"time"
)

type AuthHandler struct {
	userService *services.UserService
}

func NewAuthHandler(userService *services.UserService) *AuthHandler {
	return &AuthHandler{userService: userService}
}

func (h *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// PARSE
		var req models.UserRegisterRequest
		err := utils.BodyChecker(r, &req)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}

		// NORMALIZE
		req.Username = strings.ToLower(req.Username)
		req.Email = strings.ToLower(req.Email)
		req.Bank = strings.ToLower(req.Bank)

		// VALIDATE
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

		//UNIQUE
		errorsMap, err := h.userService.CheckUniqueness(0, &req.Username, &req.Email, &req.Phone)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if len(errorsMap) > 0 {
			b, _ := json.Marshal(errorsMap)
			utils.JSONError(w, string(b), http.StatusBadRequest)
			return
		}

		// HASHED PASSWORD
		hashedPassword, err := utils.HashedPassword(req.Password)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// CREATE USER
		user := models.User{
			Username:   req.Username,
			Email:      req.Email,
			Phone:      req.Phone,
			Password:   hashedPassword,
			Bank:       req.Bank,
			BankNumber: req.BankNumber,
			IsTenant:   req.IsTenant,
		}
		userResponse, err := h.userService.Create(&user)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// CREATE TOKEN
		token, err := utils.GenerateJWT(uint(user.ID), "user")
		if err != nil {
			utils.JSONError(w, "token generation failed", http.StatusInternalServerError)
			return
		}

		// CREATE SESSION
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

		// BUILD RESPONSE
		resp := models.UserRegisterResponse{
			User:  *userResponse,
			Token: token,
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "user registered",
			Data:    resp,
		}, http.StatusCreated)
	}
}

func (h *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// PARSE
		var req models.UserLoginRequest
		err := utils.BodyChecker(r, &req)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}

		// NORMALIZE
		req.Identifier = strings.ToLower(req.Identifier)

		// VALIDATE
		err = utils.FieldChecker(req)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}

		// FIND USER
		user, err := h.userService.Login(req.Identifier, req.Password)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// CREATE TOKEN
		token, err := utils.GenerateJWT(uint(user.ID), "user")
		if err != nil {
			utils.JSONError(w, "token generation failed", http.StatusInternalServerError)
			return
		}

		// CREATE SESSION
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

func (h *AuthHandler) Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// CHECK COOKIE
		cookie, err := r.Cookie("jwt_token_user")
		if err != nil {
			utils.JSONError(w, "no cookie", http.StatusUnauthorized)
			return
		}

		// VALIDATE
		claims, err := middleware.Validate(cookie, "user")
		if err != nil {
			utils.JSONError(w, "invalid token", http.StatusUnauthorized)
			return
		}

		// REMOVE TOKEN
		userID := int(claims["id"].(float64))
		redisKey := fmt.Sprintf("session:user:%d", userID)
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

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "logged out",
		}, http.StatusOK)
	}
}
