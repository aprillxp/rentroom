package utils

import (
	"errors"
	"rentroom/models"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func HashedPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("error hashing password")
	}
	return string(hashedPassword), nil
}

func NormalizePhone(phone string) string {
	phone = strings.TrimSpace(phone)
	phone = strings.ReplaceAll(phone, " ", "")
	phone = strings.ReplaceAll(phone, "-", "")
	phone = strings.ReplaceAll(phone, "(", "")
	phone = strings.ReplaceAll(phone, ")", "")
	return phone
}

func SeedInitialData(db *gorm.DB) {
	banks := []models.Bank{
		{Name: "BCA"},
		{Name: "Mandiri"},
		{Name: "BNI"},
	}
	for _, b := range banks {
		db.FirstOrCreate(&b, models.Bank{Name: b.Name})
	}

	countries := []models.Country{
		{Name: "Indonesia"},
		{Name: "Singapore"},
		{Name: "Malaysia"},
	}
	for _, c := range countries {
		db.FirstOrCreate(&c, models.Bank{Name: c.Name})
	}

	amenities := []models.Amenity{
		{Name: "Wifi"},
		{Name: "Parking"},
		{Name: "Pool"},
	}
	for _, a := range amenities {
		db.FirstOrCreate(&a, models.Bank{Name: a.Name})
	}
}

func PtrToStrOrEmpty(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
