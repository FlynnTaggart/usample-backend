package db

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"user-service/internal/models"
)

type PgxTx struct {
	pgx.Tx
}

func (t PgxTx) Rollback(ctx context.Context) error {
	return t.Tx.Rollback(ctx)
}
func (t PgxTx) Commit(ctx context.Context) error {
	return t.Tx.Commit(ctx)
}

func (t PgxTx) GetFriendRequestByUserIds(ctx context.Context, senderId uuid.UUID, receiverId uuid.UUID) (*models.FriendRequest, error) {
	r := &models.FriendRequest{}
	err := t.QueryRow(ctx, `select * from friend_requests where sender_id = $1 and receiver_id = $2`, senderId, receiverId).Scan(r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (t PgxTx) CreateFriendRequest(ctx context.Context, request *models.FriendRequest) error {
	var createdId uuid.UUID
	err := t.QueryRow(ctx, `insert into friend_requests (sender_id, receiver_id, is_accepted) 
	values ($1, $2, $3) returning id`,
		request.SenderId, request.ReceiverId, request.IsAccepted).Scan(&createdId)
	if err != nil {
		return err
	}
	return nil
}

func (t PgxTx) UpdateFriendRequest(ctx context.Context, request *models.FriendRequest) error {
	var updateId uuid.UUID
	err := t.QueryRow(ctx, `update friend_requests set 
                 sender_id = $2, 
                 receiver_id = $3, 
                 is_accepted = $4 
                 where id = $1 returning id`,
		request.Id,
		request.SenderId,
		request.ReceiverId,
		request.IsAccepted,
	).Scan(&updateId)
	if err != nil {
		return err
	}
	return nil
}

func (t PgxTx) DeleteFriendRequest(ctx context.Context, id uuid.UUID) error {
	var deleteId uuid.UUID
	err := t.QueryRow(ctx, `delete from friend_requests
                 where id = $1 returning id`, id).Scan(&deleteId)
	if err != nil {
		return err
	}
	return nil
}

func (t PgxTx) GetFriendRequest(ctx context.Context, id uuid.UUID) (*models.FriendRequest, error) {
	r := &models.FriendRequest{}
	err := t.QueryRow(ctx, `select * from friend_requests where id = $1`, id).Scan(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (t PgxTx) GetUserLink(ctx context.Context, id uuid.UUID) (*models.UserLink, error) {
	l := &models.UserLink{}
	err := t.QueryRow(ctx, `select * from user_links where id = $1`, id).Scan(l)
	if err != nil {
		return nil, err
	}
	return l, nil
}

func (t PgxTx) DeleteUserLink(ctx context.Context, id uuid.UUID) error {
	var deleteId uuid.UUID
	err := t.QueryRow(ctx, `delete from user_links
                 where id = $1 returning id`, id).Scan(&deleteId)
	if err != nil {
		return err
	}
	return nil
}
