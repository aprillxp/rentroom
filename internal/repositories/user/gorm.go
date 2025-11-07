package repository

import (
	"rentroom/internal/models"
	"strings"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) UserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) FindByIdentifier(identifier string) (*models.User, error) {
	var user models.User
	identifier = strings.ToLower(identifier)
	err := r.db.
		Where("username = ? OR email = ? OR phone = ?", identifier, identifier, identifier).
		First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *GormUserRepository) Create(user *models.User) error {
	user.Username = strings.ToLower(user.Username)
	user.Email = strings.ToLower(user.Email)
	return r.db.Create(user).Error
}

func (r *GormUserRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *GormUserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *GormUserRepository) UpdateFields(id uint, updates map[string]interface{}) error {
	return r.db.Model(&models.User{}).
		Where("id = ?", id).
		Updates(updates).Error
}

func (r *GormUserRepository) ExistsUsername(username string, excludeID uint) (bool, error) {
	var count int64
	query := r.db.Model(&models.User{}).
		Where("username = ?", strings.ToLower(username))
	if excludeID > 0 {
		query = query.Where("id = ?", excludeID)
	}
	err := query.Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *GormUserRepository) ExistsEmail(email string, excludeID uint) (bool, error) {
	var count int64
	query := r.db.Model(&models.User{}).
		Where("email = ?", strings.ToLower(email))
	if excludeID > 0 {
		query = query.Where("id = ?", excludeID)
	}
	err := query.Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *GormUserRepository) ExistsPhone(phone string, excludeID uint) (bool, error) {
	var count int64
	query := r.db.Model(&models.User{}).
		Where("phone = ?", strings.ToLower(phone))
	if excludeID > 0 {
		query = query.Where("id = ?", excludeID)
	}
	err := query.Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
