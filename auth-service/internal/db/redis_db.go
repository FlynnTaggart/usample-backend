package db

import (
	"auth-service/internal/models"
	"context"
	"errors"

	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
)

type RedisDB struct {
	DB *redis.Client
}

func NewRedisDB(dbUrl, dbPassword string) *RedisDB {
	return &RedisDB{
		redis.NewClient(&redis.Options{
			Addr:     dbUrl,
			Password: dbPassword,
			DB:       0,
		}),
	}
}

func (r RedisDB) CreateUserRecord(user models.User) error {
	repl := r.DB.HSet(context.Background(), user.Email, "id", user.ID.String(), "password", user.Password)
	return repl.Err()
}

func (r RedisDB) GetPasswordAndUUIDByEmail(user *models.User) error {
	repl := r.DB.HGetAll(context.Background(), user.Email)

	if repl.Err() != nil {
		return repl.Err()
	}

	values := repl.Val()
	parsedId, ok := values["id"]
	if !ok {
		return errors.New("redis db: failed to parse user id")
	}
	userId, err := uuid.Parse(parsedId)
	if err != nil {
		return err
	}
	user.ID = userId
	user.Password, ok = values["password"]
	if !ok {
		return errors.New("redis db: failed to parse user password")
	}

	return nil
}
