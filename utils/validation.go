package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"rentroom/models"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func BodyChecker(r *http.Request, req interface{}) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(req)
	if err != nil {
		return errors.New("invalid request body")
	}
	return nil
}

func FieldChecker(req interface{}) error {
	validate := validator.New()

	kind := reflect.TypeOf(req).Kind()
	if kind == reflect.Struct {
		err := validate.Struct(req)
		if err != nil {
			return ParseValidationError(err)
		}
		return nil
	}

	val := reflect.ValueOf(req)
	if val.Kind() == reflect.Ptr && val.Elem().Kind() != reflect.Struct {
		temp := struct {
			Field interface{} `validate:"required"`
		}{
			Field: req,
		}
		err := validate.Struct(temp)
		if err != nil {
			return ParseValidationError(err)
		}
		return nil
	}
	return fmt.Errorf("unsupported type for validation")
}

func ParseValidationError(err error) error {
	if errs, ok := err.(validator.ValidationErrors); ok {
		e := errs[0]
		return fmt.Errorf("field %s failed on the '%s' rule", e.Field(), e.Tag())
	}
	return err
}

func UserUniqueness(db *gorm.DB, currentUserID uint, username, email, phone string) error {
	var user models.User
	err := db.
		Where("id != ? AND (username = ? OR email = ? OR phone = ?)", currentUserID, username, email, phone).
		First(&user).Error
	if err == nil {
		return errors.New("username, email, or phone  already exists")
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return err
}

func UserIsTenant(db *gorm.DB, userID int) error {
	var user models.User
	err := db.Select("is_tenant").First(&user, userID).Error
	if err != nil {
		return err
	}
	if !user.IsTenant {
		return errors.New("user is not a tenant")
	}
	return nil
}

func PropertyUserChecker(db *gorm.DB, userID, propertyID int) error {
	var userProperty models.UserProperties
	err := db.
		Where("user_id = ? AND property_id = ?", userID, propertyID).
		First(&userProperty).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("property under tenant does not exist")
	}
	return err
}

func PasswordValidator(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}
	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return errors.New("password must contain at least one uppercase letter")
	}
	if !regexp.MustCompile(`[0-9]`).MatchString(password) {
		return errors.New("password must contain at least one number")
	}
	return nil
}

func PhoneValidator(phone string) error {
	phone = NormalizePhone(phone)
	phoneRegex := regexp.MustCompile(`^\+?[1-9]\d{6,14}$`)
	if !phoneRegex.MatchString(phone) {
		return errors.New("phone number is invalid")
	}
	return nil
}
