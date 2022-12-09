package db

import (
	"context"
	"github.com/google/uuid"
	"user-service/internal/models"
)

type DB interface {
	CreateUser(user *models.User) error
	GetUsers(limit, page int64) ([]*models.User, error)
	GetUser(id uuid.UUID) (*models.User, error)
	GetUsersByNicknamePrefix(prefix string, limit, offset int64) ([]*models.User, error)
	GetUserByNickname(nickname string) (*models.User, error)
	UpdateUserInfo(user *models.User) error
	GetUserFriends(id uuid.UUID) ([]*models.FriendRequest, error)
	GetUserSentFriends(id uuid.UUID) ([]*models.FriendRequest, error)
	GetUserReceivedFriends(id uuid.UUID) ([]*models.FriendRequest, error)
	BeginTx(ctx context.Context) (TX, error)
	AddUserLink(link *models.UserLink) error
	GetUserLinks() ([]*models.UserLink, error)
}

type TX interface {
	Rollback(ctx context.Context) error
	Commit(ctx context.Context) error
	GetFriendRequestByUserIds(ctx context.Context, senderId uuid.UUID, receiverId uuid.UUID) (*models.FriendRequest, error)
	CreateFriendRequest(ctx context.Context, request *models.FriendRequest) error
	UpdateFriendRequest(ctx context.Context, request *models.FriendRequest) error
	DeleteFriendRequest(ctx context.Context, id uuid.UUID) error
	GetFriendRequest(ctx context.Context, id uuid.UUID) (*models.FriendRequest, error)
	GetUserLink(ctx context.Context, id uuid.UUID) (*models.UserLink, error)
	DeleteUserLink(ctx context.Context, id uuid.UUID) error
}
