package services

import (
	"github.com/google/uuid"
	"user-service/internal/models"
)

func (s UsersService) CreateUser(user *models.User) error {

}

func (s UsersService) GetUsers(limit, offset int64) ([]*models.User, error) {

}

func (s UsersService) GetUser(id uuid.UUID) (*models.User, error) {

}

func (s UsersService) GetUsersByNicknamePrefix(prefix string, limit int64, offset int64) ([]*models.User, error) {

}

func (s UsersService) GetUserByNickname(nickname string) (models.User, error) {

}

func (s UsersService) UpdateUserInfo(user *models.User) error {

}
