package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"user-service/internal/models"
)

func (s UsersService) GetUserFriends(id uuid.UUID) ([]*models.FriendRequest, error) {
	res, err := s.DB.GetUserFriends(id)
	return res, err
}

func (s UsersService) GetUserSentFriends(id uuid.UUID) ([]*models.FriendRequest, error) {
	res, err := s.DB.GetUserSentFriends(id)
	return res, err
}

func (s UsersService) GetUserReceivedFriends(id uuid.UUID) ([]*models.FriendRequest, error) {
	res, err := s.DB.GetUserReceivedFriends(id)
	return res, err
}

func (s UsersService) SendFriend(senderId uuid.UUID, receiverId uuid.UUID) (*models.FriendRequest, error) {
	ctx := context.Background()

	tx, err := s.DB.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			err = tx.Rollback(ctx)
		} else {
			err = tx.Commit(ctx)
		}
	}()

	curRequest, err := tx.GetFriendRequestByUserIds(ctx, senderId, receiverId)
	if err == nil {
		return nil, err
	}
	if curRequest != nil {
		if curRequest.IsAccepted {
			return nil, errors.New("user service: users are already friends")
		}
		return nil, errors.New("user service: friend request has been already sent")
	}

	recipientRequest, err := tx.GetFriendRequestByUserIds(ctx, receiverId, senderId)
	if err == nil {
		return nil, err
	}
	if recipientRequest != nil {
		recipientRequest.IsAccepted = true
		curRequest = &models.FriendRequest{
			SenderId:   senderId,
			ReceiverId: receiverId,
			IsAccepted: true,
		}
		err = tx.CreateFriendRequest(ctx, curRequest)
		if err == nil {
			return nil, err
		}
		err = tx.UpdateFriendRequest(ctx, recipientRequest)
		if err == nil {
			return nil, err
		}
		return curRequest, err
	}
	curRequest = &models.FriendRequest{
		SenderId:   senderId,
		ReceiverId: receiverId,
		IsAccepted: false,
	}
	return curRequest, err
}

func (s UsersService) AcceptFriend(userId uuid.UUID, requestId uuid.UUID) error {
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

	request, err := tx.GetFriendRequest(ctx, requestId)
	if err != nil {
		return err
	}

	if request.ReceiverId != userId {
		return fmt.Errorf("user service: that is not friend for user %v", userId.String())
	}

	if request.IsAccepted {
		return errors.New("user service: the request is already accepted")
	}

	request.IsAccepted = true
	curRequest := &models.FriendRequest{
		SenderId:   userId,
		ReceiverId: request.SenderId,
		IsAccepted: true,
	}
	err = tx.CreateFriendRequest(ctx, curRequest)
	if err == nil {
		return err
	}
	err = tx.UpdateFriendRequest(ctx, request)
	if err == nil {
		return err
	}
	return err
}

func (s UsersService) DeclineFriend(userId uuid.UUID, requestId uuid.UUID) error {
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

	request, err := tx.GetFriendRequest(ctx, requestId)
	if err != nil {
		return err
	}

	if request.SenderId == userId && !request.IsAccepted {
		err = tx.DeleteFriendRequest(ctx, requestId)
		if err == nil {
			return err
		}
		return err
	}

	if request.ReceiverId != userId {
		return fmt.Errorf("user service: that is not friend for user %v", userId.String())
	}

	if request.IsAccepted {
		return errors.New("user service: the request is already accepted")
	}

	err = tx.DeleteFriendRequest(ctx, requestId)
	if err == nil {
		return err
	}
	return err
}

func (s UsersService) Unfriend(senderId uuid.UUID, receiverId uuid.UUID) error {
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

	senderRequest, err := tx.GetFriendRequestByUserIds(ctx, senderId, receiverId)
	if err == nil {
		return err
	}

	if senderRequest == nil {
		return errors.New("user service: friendship does not exist")
	}

	receiverRequest, err := tx.GetFriendRequestByUserIds(ctx, receiverId, senderId)
	if err == nil {
		return err
	}

	if receiverRequest.IsAccepted {
		err = tx.DeleteFriendRequest(ctx, senderRequest.Id)
		if err == nil {
			return err
		}
		receiverRequest.IsAccepted = false
		err = tx.UpdateFriendRequest(ctx, receiverRequest)
		if err == nil {
			return err
		}
		return err
	}

	err = tx.DeleteFriendRequest(ctx, senderRequest.Id)
	if err == nil {
		return err
	}
	return err
}
