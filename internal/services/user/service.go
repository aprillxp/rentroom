package service

import (
	"errors"
	"rentroom/internal/models"
	repository "rentroom/internal/repositories/user"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetByID(id uint) (*models.UserResponse, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return NewUserResponse(user), err
}

func (s *UserService) CheckUniqueness(excludeID uint, username, email, phone *string) (map[string]string, error) {
	errorsMap := make(map[string]string)

	if username != nil {
		exists, err := s.repo.ExistsUsername(*username, excludeID)
		if err != nil {
			return nil, err
		}
		if exists {
			errorsMap["username"] = "username already taken"
		}
	}
	if email != nil {
		exists, err := s.repo.ExistsEmail(*email, excludeID)
		if err != nil {
			return nil, err
		}
		if exists {
			errorsMap["email"] = "email already registered"
		}
	}
	if phone != nil {
		exists, err := s.repo.ExistsPhone(*phone, excludeID)
		if err != nil {
			return nil, err
		}
		if exists {
			errorsMap["phone"] = "phone number already used"
		}
	}

	return errorsMap, nil
}

func (s *UserService) Create(user *models.User) (*models.UserResponse, error) {
	err := s.repo.Create(user)
	if err != nil {
		return nil, errors.New("failed to create user")
	}
	return NewUserResponse(user), nil
}

func (s *UserService) Update(id uint, updates map[string]interface{}) (*models.UserResponse, error) {
	err := s.repo.UpdateFields(id, updates)
	if err != nil {
		return nil, err
	}
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return NewUserResponse(user), err
}

func (s *UserService) Login(identifier, password string) (*models.UserResponse, error) {
	user, err := s.repo.FindByIdentifier(identifier)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return NewUserResponse(user), nil
}
