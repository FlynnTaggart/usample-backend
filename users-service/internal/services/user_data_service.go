package services

import (
	"github.com/google/uuid"
	"user-service/internal/models"
)

func (s UsersService) CreateUser(user *models.User) error {
	return s.DB.CreateUser(user)
}

func (s UsersService) GetUsers(limit, offset int64) ([]*models.User, error) {
	res, err := s.DB.GetUsers(limit, offset)
	return res, err
}

func (s UsersService) GetUser(id uuid.UUID) (*models.User, error) {
	res, err := s.DB.GetUser(id)
	return res, err
}

func (s UsersService) GetUsersByNicknamePrefix(prefix string, limit int64, offset int64) ([]*models.User, error) {
	res, err := s.DB.GetUsersByNicknamePrefix(prefix, limit, offset)
	return res, err
}

func (s UsersService) GetUserByNickname(nickname string) (*models.User, error) {
	res, err := s.DB.GetUserByNickname(nickname)
	return res, err
}

func (s UsersService) UpdateUserInfo(user *models.User) error {
	return s.DB.UpdateUserInfo(user)
}
