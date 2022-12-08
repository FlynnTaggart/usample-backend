package services

import (
	"auth-service/internal/db"
	"auth-service/internal/models"
	"auth-service/utils"

	"errors"

	"github.com/google/uuid"
)

type AuthService struct {
	DB  db.DB
	JWT utils.JwtWrapper
}

func NewAuthService(DB db.DB, JWT utils.JwtWrapper) *AuthService {
	return &AuthService{DB: DB, JWT: JWT}
}

func (s AuthService) Register(email, password string) (string, error) {
	var user models.User

	err := s.DB.GetPasswordAndUUIDByEmail(&models.User{Email: email})

	if err == nil {
		return "", errors.New("auth service: user with such email already exists")
	}

	user.Email = email
	user.Password = utils.HashPassword(password)
	user.ID = uuid.New()

	err = s.DB.CreateUserRecord(user)

	if err != nil {
		return "", err
	}

	return user.ID.String(), nil
}

func (s AuthService) Login(email, password string) (string, error) {
	user := &models.User{
		Email: email,
	}

	err := s.DB.GetPasswordAndUUIDByEmail(user)

	if err != nil {
		return "", err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("auth service: wrong password")
	}

	token, err := s.JWT.GenerateToken(user)

	if err != nil {
		return "", errors.New("auth service: failed to generate token")
	}

	return token, nil
}

func (s AuthService) Validate(token string) (uuid.UUID, error) {
	claims, err := s.JWT.ValidateToken(token)

	if err != nil {
		return uuid.Nil, err
	}

	user := &models.User{
		Email: claims.Email,
	}

	err = s.DB.GetPasswordAndUUIDByEmail(user)

	if err != nil {
		return uuid.Nil, err
	}

	return user.ID, nil
}
