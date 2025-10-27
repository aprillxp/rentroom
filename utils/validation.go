package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"rentroom/internal/models"
	"time"

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

// USER
func UserUniqueness(db *gorm.DB, currentUserID uint, username, email, phone string) error {
	var user models.User
	err := db.
		Where("id != ? AND (username = ? OR email = ? OR phone = ?)", currentUserID, username, email, phone).
		First(&user).Error
	if err == nil {
		return errors.New("username, email, or phone already exists")
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return err
}
func UserIsTenant(db *gorm.DB, userID uint) error {
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

// PROPERTIES
func PropertyExist(db *gorm.DB, propertyID uint) error {
	var property models.Property
	err := db.First(&property, propertyID).Error
	if err != nil {
		return errors.New("property not found")
	}
	return nil
}
func PropertyUserChecker(db *gorm.DB, userID, propertyID uint) error {
	var userProperty models.UserProperties
	err := db.
		Where("user_id = ? AND property_id = ?", userID, propertyID).
		First(&userProperty).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("property under tenant does not exist")
	}
	return err
}
func PropertOwnedByUser(db *gorm.DB, userID, propertyID uint) error {
	var userProperty models.UserProperties
	err := db.
		Where("user_id = ? AND property_id = ?", userID, propertyID).
		First(&userProperty).Error
	if err == nil {
		return errors.New("cannot perform action on your own property")
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return err
}
func PropertyAvailable(db *gorm.DB, propertyID uint, checkin, checkout time.Time) error {
	var property models.Property
	err := db.First(&property, propertyID).Error
	if err != nil {
		return err
	}
	if (checkin.Equal(property.DisabledDateTo) || checkin.Before(property.DisabledDateTo)) &&
		(checkout.Equal(property.DisabledDateFrom) || checkout.After(property.DisabledDateFrom)) {
		return errors.New("property is unavailable on your date request by tennant")
	}
	if property.Status != models.StatusPublished {
		return errors.New("property is not published")
	}
	var transactions = []models.Transaction{}
	err = db.
		Where("property_id = ? AND status = ?", propertyID, models.StatusApproved).
		Find(&transactions).Error
	if err != nil {
		return err
	}
	for _, t := range transactions {
		if checkin.Before(t.CheckOut) && t.CheckIn.Before(checkout) {
			return errors.New("property is already booked for your requested dates")
		}
	}
	return nil
}
func PropertyHaveAnActiveTransaction(db *gorm.DB, propertyId uint) error {
	var propertyTransaction models.Transaction
	err := db.
		Where("property_id = ? AND status = ?", propertyId, models.StatusApproved).
		First(&propertyTransaction).Error
	if err == nil {
		return errors.New("cannot perform action, property has an active tranacstion")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}

// TRANSACTION
func TransactionExist(db *gorm.DB, tranacstionID uint) error {
	var transaction models.Transaction
	err := db.First(&transaction, tranacstionID).Error
	if err != nil {
		return errors.New("transaction not found")
	}
	return nil
}
func TransactionUserChecker(db *gorm.DB, userID uint, transactionID uint) error {
	var userTrasaction models.Transaction
	err := db.
		Where("id = ? AND user_id = ?", transactionID, userID).
		First(&userTrasaction).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("transaction under user does not exist")
	}
	return err
}
func TransactionOwnedByUser(db *gorm.DB, userID, propertyID uint) error {
	var transaction models.Transaction
	err := db.
		Where("user_id = ? AND property_id = ? AND status IN ?", userID, propertyID, []int{models.StatusPending, models.StatusApproved}).
		First(&transaction).Error
	if err == nil {
		return errors.New("transaction under this property already created")
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return err
}
func TransactionIsPending(db *gorm.DB, transactionID uint) error {
	var transaction models.Transaction
	err := db.
		Where("id = ? AND status = ?", transactionID, models.StatusPending).
		First(&transaction).Error
	if err == nil {
		return nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("cannot perform action since status is not pending")
	}
	return nil
}
func TransactionIsApproved(db *gorm.DB, transactionID uint) error {
	var transaction models.Transaction
	err := db.
		Where("id = ? AND status = ?", transactionID, models.StatusApproved).
		First(&transaction).Error
	if err == nil {
		return nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("cannot perform action since satus is not approved")
	}
	return nil
}
func TransactionIsDone(db *gorm.DB, transactionID uint) error {
	var transaction models.Transaction
	err := db.
		Where("id = ? AND status = ?", transactionID, models.StatusDone).
		First(&transaction).Error
	if err == nil {
		return nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("cannot perform action since satus is not done")
	}
	return nil
}

// GENERAL
func BankValidator(db *gorm.DB, bankID int) error {
	var bank models.Bank
	err := db.
		Select("id").
		First(&bank, bankID).Error
	if err != nil {
		return errors.New("bank id is not found")
	}
	return nil
}
func CountryValidator(db *gorm.DB, CountryID uint) error {
	var country models.Country
	err := db.
		Select("id").
		First(&country, CountryID).Error
	if err != nil {
		return errors.New("country id is not found")
	}
	return nil
}
func VoucherUniqueness(db *gorm.DB, name string) error {
	var voucher models.Voucher
	err := db.
		Where("name = ?", name).
		First(&voucher).Error
	if err == nil {
		return errors.New("voucher name already exists")
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return err
}
func ReviewUniqueness(db *gorm.DB, tranacstionID uint) error {
	var review models.Review
	err := db.
		Where("transaction_id = ?", tranacstionID).
		First(&review).Error
	if err == nil {
		return errors.New("review name already exists")
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return err
}
func CountryUniqueness(db *gorm.DB, countryName string) error {
	var country models.Country
	err := db.
		Where("name = ?", countryName).
		First(&country).Error
	if err == nil {
		return errors.New("country name already exists")
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return err
}
func CountryHaveProperty(db *gorm.DB, countryID uint) error {
	var countryProperty models.Property
	err := db.
		Where("country_id = ?", countryID,).
		First(&countryProperty).Error
	if err == nil {
		return errors.New("cannot perform action, country has an property")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}
