package users

import (
	"net/http"
	"rentroom/middleware"
	"rentroom/utils"

	"gorm.io/gorm"
)

func UserGet(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// AUTH
		userID, err := middleware.MustUserID(r)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// QUERY
		user, err := utils.GetUser(db, userID)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "user returned successfully",
			Data:    user,
		}, http.StatusOK)
	}
}
