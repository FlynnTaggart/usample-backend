package services

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"user-service/internal/models"
)

func (s UsersService) AddUserLink(link *models.UserLink) error {
	return s.DB.AddUserLink(link)
}

func (s UsersService) GetUserLinks() ([]*models.UserLink, error) {
	res, err := s.DB.GetUserLinks()
	return res, err
}

func (s UsersService) DeleteUserLink(userId uuid.UUID, id uuid.UUID) error {
	ctx := context.Background()

	tx, err := s.DB.BeginTx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			err = tx.Rollback(ctx)
		} else {
			err = tx.Commit(ctx)
		}
	}()

	link, err := tx.GetUserLink(ctx, id)
	if err != nil {
		return err
	}

	if link.UserId != userId {
		return errors.New("user service: user id does not match with link user id")
	}

	err = tx.DeleteUserLink(ctx, id)
	return err
}
