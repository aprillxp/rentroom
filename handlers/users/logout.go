package users

import (
	"fmt"
	"net/http"
	"rentroom/middleware"
	"rentroom/utils"
)

func UserLogout() http.HandlerFunc {
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
