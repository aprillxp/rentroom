package utils

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"rentroom/models"
	"time"

	"gorm.io/gorm"
)

func GetProperty(db *gorm.DB, propertyID int) (models.Property, error) {
	var property models.Property
	err := db.First(&property, propertyID).Error
	if err != nil {
		return property, errors.New("property not found")
	}
	return property, nil
}

func GetUser(db *gorm.DB, userID uint) (models.UserResponse, error) {
	var user models.User
	err := db.First(&user, userID).Error
	if err != nil {
		return models.UserResponse{}, errors.New("user not found")
	}
	return models.UserResponse{
		ID:         user.ID,
		Username:   user.Username,
		Email:      user.Email,
		Phone:      user.Phone,
		BankID:     user.BankID,
		BankNumber: user.BankNumber,
		IsTenant:   user.IsTenant,
	}, nil
}

func GetVoucher(db *gorm.DB, voucherID int) (models.Voucher, error) {
	var voucher models.Voucher
	err := db.First(&voucher, voucherID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Voucher{}, errors.New("voucher not found")
		}
		return models.Voucher{}, errors.New("voucher not found")
	}
	return voucher, err
}

func GetTransaction(db *gorm.DB, transactionID uint) (models.TransactionResponse, error) {
	var transaction models.Transaction
	err := db.
		Where("id = ?", transactionID).
		First(&transaction).Error
	if err != nil {
		return models.TransactionResponse{}, errors.New("transaction not found")
	}
	return models.TransactionResponse{
		ID:         transaction.ID,
		PropertyID: transaction.PropertyID,
		Price:      transaction.Price,
		CheckIn:    transaction.CheckIn,
		CheckOut:   transaction.CheckOut,
		Status:     transaction.Status,
		VoucherID:  transaction.VoucherID,
	}, nil
}

func GetUserTransaction(db *gorm.DB, userID, transactionID uint) (models.TransactionResponse, error) {
	var transaction models.Transaction
	err := db.
		Where("id = ? AND user_id = ?", transactionID, userID).
		First(&transaction).Error
	if err != nil {
		return models.TransactionResponse{}, errors.New("transaction not found")
	}
	return models.TransactionResponse{
		ID:         transaction.ID,
		PropertyID: transaction.PropertyID,
		Price:      transaction.Price,
		CheckIn:    transaction.CheckIn,
		CheckOut:   transaction.CheckOut,
		Status:     transaction.Status,
		VoucherID:  transaction.VoucherID,
	}, nil
}

func GetTenantTransaction(db *gorm.DB, propertyIDs []uint, transactionID uint) (models.TransactionResponse, error) {
	var transaction models.Transaction
	err := db.
		Where("id = ? AND property_id IN ?", transactionID, propertyIDs).
		First(&transaction).Error
	if err != nil {
		return models.TransactionResponse{}, errors.New("transaction not found")
	}
	return models.TransactionResponse{
		ID:         transaction.ID,
		PropertyID: transaction.PropertyID,
		Price:      transaction.Price,
		CheckIn:    transaction.CheckIn,
		CheckOut:   transaction.CheckOut,
		Status:     transaction.Status,
		VoucherID:  transaction.VoucherID,
	}, nil
}

func GetPropertyIDs(db *gorm.DB, userID uint) ([]uint, error) {
	var propertyIDs []uint
	err := db.Model(&models.UserProperties{}).
		Where("user_id = ?", userID).
		Pluck("property_id", &propertyIDs).Error
	if err != nil {
		return nil, err
	}
	if len(propertyIDs) == 0 {
		return nil, errors.New("no properties found for this user")
	}
	return propertyIDs, nil
}

func UploadImagesProperty(tx *gorm.DB, r *http.Request, propertyID uint) ([]models.Image, error) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return nil, errors.New("failed to parse form")
	}
	files := r.MultipartForm.File["images"]
	var uploaded []models.Image

	for _, header := range files {
		file, err := header.Open()
		if err != nil {
			return nil, errors.New("failed to open file")
		}
		defer file.Close()

		filename := fmt.Sprintf("%d-%s", time.Now().Unix(), header.Filename)
		fsPath := "./uploads/" + filename
		publicPath := "/uploads/" + filename

		out, err := os.Create(fsPath)
		if err != nil {
			return nil, errors.New("failed to save file")
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			return nil, errors.New("failed to write file")
		}

		img := models.Image{
			PropertyID: uint(propertyID),
			Path:       publicPath,
		}
		err = tx.Create(&img).Error
		if err != nil {
			_ = os.Remove(fsPath)
			return nil, errors.New("failed to save image record")
		}

		uploaded = append(uploaded, img)
	}

	return uploaded, nil
}
