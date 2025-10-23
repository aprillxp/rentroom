package middleware

import (
	"context"
	"fmt"
	"net/http"
	"rentroom/utils"
)

func JwtAuthUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("jwt_token_user")
		if err != nil {
			utils.JSONError(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		claims, err := Validate(c, "user")
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}

		userID := int(claims["id"].(float64))
		redisKey := fmt.Sprintf("session:user:%d", userID)

		storedToken, err := utils.RedisUser.Get(utils.Ctx, redisKey).Result()
		if err != nil || storedToken != c.Value {
			utils.JSONError(w, "session expired or invalid", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), CtxUserID, userID)
		ctx = context.WithValue(ctx, CtxRole, "user")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func JwtAuthAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("jwt_token_admin")
		if err != nil {
			utils.JSONError(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		claims, err := Validate(c, "admin")
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}

		adminID := int(claims["id"].(float64))
		redisKey := fmt.Sprintf("session:admin:%d", adminID)

		storedToken, err := utils.RedisUser.Get(utils.Ctx, redisKey).Result()
		if err != nil || storedToken != c.Value {
			utils.JSONError(w, "session expired or invalid", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), CtxAdminID, adminID)
		ctx = context.WithValue(ctx, CtxRole, "user")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
