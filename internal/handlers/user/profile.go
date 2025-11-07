package user

import (
	"encoding/json"
	"net/http"
	"rentroom/internal/models"
	services "rentroom/internal/services/user"
	"rentroom/middleware"
	"rentroom/utils"
	"strings"
)

type ProfileHandler struct {
	userService *services.UserService
}

func NewProfileHandler(userService *services.UserService) *ProfileHandler {
	return &ProfileHandler{userService: userService}
}

func (h *ProfileHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// AUTH
		userID, err := middleware.MustUserID(r)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// QUERY
		user, err := h.userService.GetByID(userID)
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

func (h *ProfileHandler) Edit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// AUTH
		userID, err := middleware.MustUserID(r)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// PARSE
		var req models.UserEditRequest
		err = utils.BodyChecker(r, &req)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}

		// NORMALIZE
		if req.Username != nil {
			*req.Username = strings.ToLower(*req.Username)
		}
		if req.Email != nil {
			*req.Email = strings.ToLower(*req.Email)
		}

		// VALIDATE
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

		//UNIQUE
		errorsMap, err := h.userService.CheckUniqueness(userID, req.Username, req.Email, req.Phone)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if len(errorsMap) > 0 {
			b, _ := json.Marshal(errorsMap)
			utils.JSONError(w, string(b), http.StatusBadRequest)
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
		if req.Bank != nil {
			updates["bank_id"] = *req.Bank
		}
		if req.BankNumber != nil {
			updates["bank_number"] = *req.BankNumber
		}
		userUpdated, err := h.userService.Update(userID, updates)
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
