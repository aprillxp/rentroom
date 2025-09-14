package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"rentroom/models"

	"gorm.io/gorm"
)

func BodyChecker(r *http.Request, req interface{}) error {
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return errors.New("invalid request")
	}
	return nil
}

func UserUniqueness(db *gorm.DB, username, email, phone string) error {
	var user models.User
	err := db.
		Where("username = ? OR email = ? OR phone = ?", username, email, phone).
		First(&user).Error
	if err == nil {
		return errors.New("username, email, or phone  already exists")
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return err
}
