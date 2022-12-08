package db

import (
	"auth-service/internal/models"
)

type DB interface {
	CreateUserRecord(user models.User) error
	GetPasswordAndUUIDByEmail(user *models.User) error
}
