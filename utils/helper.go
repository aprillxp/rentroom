package utils

import (
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"rentroom/models"
	"strings"
	"time"

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

func PathImage(header *multipart.FileHeader) (fsPath, publicPath string) {
	filename := fmt.Sprintf("%d-%s", time.Now().Unix(), header.Filename)
	fsPath = filepath.Join("./uploads", filename)
	publicPath = "/uploads/" + filename
	return
}

func ConvertPropertiesResponse(properties []models.Property) []models.PropertyResponse {
	responses := make([]models.PropertyResponse, 0, len(properties))
	for _, p := range properties {
		responses = append(responses, models.PropertyResponse{
			ID:               p.ID,
			CountryID:        p.CountryID,
			Name:             p.Name,
			Guests:           p.Guests,
			Price:            p.Price,
			Status:           p.Status,
			DisabledDateFrom: p.DisabledDateFrom,
			DisabledDateTo:   p.DisabledDateTo,
			Description:      p.Description,
			Geo:              p.Geo,
			Province:         p.Province,
			District:         p.District,
			City:             p.City,
			Address:          p.Address,
			Zip:              p.Zip,
		})
	}
	return responses
}

func ConvertTransactionsResponse(transactions []models.Transaction) []models.TransactionResponse {
	responses := make([]models.TransactionResponse, 0, len(transactions))
	for _, t := range transactions {
		responses = append(responses, models.TransactionResponse{
			ID:         t.ID,
			UserID:     t.ID,
			PropertyID: t.PropertyID,
			Price:      t.Price,
			CheckIn:    t.CheckIn,
			CheckOut:   t.CheckOut,
			Status:     t.Status,
			VoucherID:  t.VoucherID,
		})
	}
	return responses
}

func ConvertTransactionResponse(transaction models.Transaction) models.TransactionResponse {
	return models.TransactionResponse{
		ID:         transaction.ID,
		PropertyID: transaction.PropertyID,
		Price:      transaction.Price,
		CheckIn:    transaction.CheckIn,
		CheckOut:   transaction.CheckOut,
		Status:     transaction.Status,
		VoucherID:  transaction.VoucherID,
	}
}
