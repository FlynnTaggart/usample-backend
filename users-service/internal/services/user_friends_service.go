package services

import (
	"github.com/google/uuid"
	"user-service/internal/models"
)

func (s UsersService) GetUserFriends(id uuid.UUID) ([]*models.FriendRequest, error) {

}

func (s UsersService) GetUserSentFriends(id uuid.UUID) ([]*models.FriendRequest, error) {

}

func (s UsersService) GetUserReceivedFriends(id uuid.UUID) ([]*models.FriendRequest, error) {

}

func (s UsersService) SendFriend(senderId uuid.UUID, receiverId uuid.UUID) (*models.FriendRequest, error) {

}

func (s UsersService) AcceptFriend(userId uuid.UUID, requestId uuid.UUID) error {

}

func (s UsersService) DeclineFriend(userId uuid.UUID, requestId uuid.UUID) error {

}

func (s UsersService) Unfriend(senderId uuid.UUID, receiverId uuid.UUID) error {

}
