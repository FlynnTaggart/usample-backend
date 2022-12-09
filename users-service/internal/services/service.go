package services

import (
	"user-service/internal/db"
	"user-service/pkg/logger"
)

type UsersService struct {
	DB  db.DB
	log logger.Logger
}

func NewUsersService(DB db.DB, log logger.Logger) *UsersService {
	return &UsersService{DB: DB, log: log}
}
