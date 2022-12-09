package db

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	"user-service/internal/models"
)

type PgxDB struct {
	*pgxpool.Pool
	Logger tracelog.Logger
}

func NewPgxDB(pool *pgxpool.Pool, logger tracelog.Logger) *PgxDB {
	return &PgxDB{
		Pool:   pool,
		Logger: logger,
	}
}

func (p PgxDB) BeginTx(ctx context.Context) (TX, error) {
	tx, err := p.Pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.Serializable,
	})
	if err != nil {
		return nil, err
	}

	return PgxTx{
		Tx: tx,
	}, nil
}

func (p PgxDB) CreateUser(user *models.User) error {
	ctx := context.Background()
	var createdId uuid.UUID
	err := p.QueryRow(ctx, `insert into users (id, nickname, first_name, second_name, default_access_type, user_type, bio) 
	values ($1, $2, $3, $4, $5, $6, $7) returning id`,
		user.Id,
		user.Nickname,
		user.FirstName,
		user.SecondName,
		user.DefaultAccessType,
		user.UserType,
		user.Bio).Scan(&createdId)
	if err != nil {
		return err
	}
	return nil
}

func (p PgxDB) GetUsers(limit, page int64) ([]*models.User, error) {
	ctx := context.Background()

	offset := limit * (page - 1)
	rows, err := p.Query(ctx, `select * from users order by nickname limit $1 offset $2`, limit, offset)
	if err != nil {
		return nil, err
	}
	var res []*models.User

	for rows.Next() {
		u := &models.User{}
		err = rows.Scan(
			&u.Id,
			&u.Nickname,
			&u.FirstName,
			&u.SecondName,
			&u.DefaultAccessType,
			&u.UserType,
			&u.Bio)
		if err != nil {
			return nil, err
		}
		res = append(res, u)
	}

	return res, nil
}

func (p PgxDB) GetUser(id uuid.UUID) (*models.User, error) {
	ctx := context.Background()

	u := &models.User{}
	err := p.QueryRow(ctx, `select * from users where id = $1`, id).Scan(
		&u.Id,
		&u.Nickname,
		&u.FirstName,
		&u.SecondName,
		&u.DefaultAccessType,
		&u.UserType,
		&u.Bio)
	if err != nil {
		return nil, err
	}

	p.Logger.Log(ctx, tracelog.LogLevelError, u.Nickname, map[string]interface{}{})
	return u, nil
}

func (p PgxDB) GetUsersByNicknamePrefix(prefix string, limit, page int64) ([]*models.User, error) {
	ctx := context.Background()

	offset := limit * (page - 1)
	rows, err := p.Query(ctx, `select * from users where nickname ilike '$3%' order by nickname limit $1 offset $2`, limit, offset, prefix)
	if err != nil {
		return nil, err
	}
	var res []*models.User

	for rows.Next() {
		u := &models.User{}
		err = rows.Scan(
			&u.Id,
			&u.Nickname,
			&u.FirstName,
			&u.SecondName,
			&u.DefaultAccessType,
			&u.UserType,
			&u.Bio)
		if err != nil {
			return nil, err
		}
		res = append(res, u)
	}

	return res, nil
}

func (p PgxDB) GetUserByNickname(nickname string) (*models.User, error) {
	ctx := context.Background()

	u := &models.User{}
	err := p.QueryRow(ctx, `select * from users where nickname = $1`, nickname).Scan(
		&u.Id,
		&u.Nickname,
		&u.FirstName,
		&u.SecondName,
		&u.DefaultAccessType,
		&u.UserType,
		&u.Bio)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (p PgxDB) UpdateUserInfo(user *models.User) error {
	ctx := context.Background()
	var updateId uuid.UUID
	err := p.QueryRow(ctx, `update users set 
                 nickname = $2, 
                 first_name = $3, 
                 second_name = $4, 
                 default_access_type = $5, 
                 user_type = $6, 
                 bio = $7 where id = $1 returning id`,
		user.Id,
		user.Nickname,
		user.FirstName,
		user.SecondName,
		user.DefaultAccessType,
		user.UserType,
		user.Bio).Scan(&updateId)
	if err != nil {
		return err
	}
	return nil
}

func (p PgxDB) GetUserFriends(id uuid.UUID) ([]*models.FriendRequest, error) {
	ctx := context.Background()

	rows, err := p.Query(ctx, `select * from friend_requests where sender_id = $1 and is_accepted = true`, id)
	if err != nil {
		return nil, err
	}
	var res []*models.FriendRequest

	for rows.Next() {
		f := &models.FriendRequest{}
		err = rows.Scan(&f.Id, &f.SenderId, &f.ReceiverId, &f.IsAccepted)
		if err != nil {
			return nil, err
		}
		res = append(res, f)
	}

	return res, nil
}

func (p PgxDB) GetUserSentFriends(id uuid.UUID) ([]*models.FriendRequest, error) {
	ctx := context.Background()

	rows, err := p.Query(ctx, `select * from friend_requests where sender_id = $1 and is_accepted = false`, id)
	if err != nil {
		return nil, err
	}
	var res []*models.FriendRequest

	for rows.Next() {
		f := &models.FriendRequest{}
		err = rows.Scan(&f.Id, &f.SenderId, &f.ReceiverId, &f.IsAccepted)
		if err != nil {
			return nil, err
		}
		res = append(res, f)
	}

	return res, nil
}

func (p PgxDB) GetUserReceivedFriends(id uuid.UUID) ([]*models.FriendRequest, error) {
	ctx := context.Background()

	rows, err := p.Query(ctx, `select * from friend_requests where receiver_id = $1 and is_accepted = false`, id)
	if err != nil {
		return nil, err
	}
	var res []*models.FriendRequest

	for rows.Next() {
		f := &models.FriendRequest{}
		err = rows.Scan(&f.Id, &f.SenderId, &f.ReceiverId, &f.IsAccepted)
		if err != nil {
			return nil, err
		}
		res = append(res, f)
	}

	return res, nil
}

func (p PgxDB) AddUserLink(link *models.UserLink) error {
	ctx := context.Background()
	var createdId uuid.UUID
	err := p.QueryRow(ctx, `insert into user_links (link_type, link_url, user_id) 
	values ($1, $2, $3) returning id`,
		link.Type,
		link.Url,
		link.UserId).Scan(&createdId)
	if err != nil {
		return err
	}
	return nil
}

func (p PgxDB) GetUserLinks() ([]*models.UserLink, error) {
	ctx := context.Background()

	rows, err := p.Query(ctx, `select * from user_links order by link_type, link_url`)
	if err != nil {
		return nil, err
	}
	var res []*models.UserLink

	for rows.Next() {
		l := &models.UserLink{}
		err = rows.Scan(&l.Id, &l.Type, &l.Url, &l.UserId)
		if err != nil {
			return nil, err
		}
		res = append(res, l)
	}

	return res, nil
}
