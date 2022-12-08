package servers

import (
	"context"
	"github.com/google/uuid"
	"net/http"
	"user-service/internal/pb"
)

func (s UsersServer) GetUserFriends(_ context.Context, req *pb.GetUserFriendsRequest) (*pb.GetUserFriendsResponse, error) {
	id, err := uuid.Parse(req.UserId)
	if err != nil {
		return &pb.GetUserFriendsResponse{
			Status: http.StatusBadRequest,
			Error:  "user service: invalid user id",
		}, nil
	}

	res, err := s.service.GetUserFriends(id)
	if err != nil {
		return &pb.GetUserFriendsResponse{
			Status: http.StatusBadGateway,
			Error:  err.Error(),
		}, nil
	}

	resp := &pb.GetUserFriendsResponse{
		Status:         http.StatusOK,
		FriendRequests: []*pb.FriendRequest{},
	}

	for _, v := range res {
		resp.FriendRequests = append(resp.FriendRequests, &pb.FriendRequest{
			Id:         v.Id.String(),
			SenderId:   v.SenderId.String(),
			ReceiverId: v.ReceiverId.String(),
			IsAccepted: v.IsAccepted,
		})
	}

	return resp, err
}

func (s UsersServer) GetUserSentFriends(_ context.Context, req *pb.GetUserFriendsRequest) (*pb.GetUserFriendsResponse, error) {
	id, err := uuid.Parse(req.UserId)
	if err != nil {
		return &pb.GetUserFriendsResponse{
			Status: http.StatusBadRequest,
			Error:  "user service: invalid user id",
		}, nil
	}

	res, err := s.service.GetUserSentFriends(id)
	if err != nil {
		return &pb.GetUserFriendsResponse{
			Status: http.StatusBadGateway,
			Error:  err.Error(),
		}, nil
	}

	resp := &pb.GetUserFriendsResponse{
		Status:         http.StatusOK,
		FriendRequests: []*pb.FriendRequest{},
	}

	for _, v := range res {
		resp.FriendRequests = append(resp.FriendRequests, &pb.FriendRequest{
			Id:         v.Id.String(),
			SenderId:   v.SenderId.String(),
			ReceiverId: v.ReceiverId.String(),
			IsAccepted: v.IsAccepted,
		})
	}

	return resp, err
}

func (s UsersServer) GetUserReceivedFriends(_ context.Context, req *pb.GetUserFriendsRequest) (*pb.GetUserFriendsResponse, error) {
	id, err := uuid.Parse(req.UserId)
	if err != nil {
		return &pb.GetUserFriendsResponse{
			Status: http.StatusBadRequest,
			Error:  "user service: invalid user id",
		}, nil
	}

	res, err := s.service.GetUserReceivedFriends(id)
	if err != nil {
		return &pb.GetUserFriendsResponse{
			Status: http.StatusBadGateway,
			Error:  err.Error(),
		}, nil
	}

	resp := &pb.GetUserFriendsResponse{
		Status:         http.StatusOK,
		FriendRequests: []*pb.FriendRequest{},
	}

	for _, v := range res {
		resp.FriendRequests = append(resp.FriendRequests, &pb.FriendRequest{
			Id:         v.Id.String(),
			SenderId:   v.SenderId.String(),
			ReceiverId: v.ReceiverId.String(),
			IsAccepted: v.IsAccepted,
		})
	}

	return resp, err
}

func (s UsersServer) SendFriend(_ context.Context, req *pb.SendFriendRequest) (*pb.SendFriendResponse, error) {
	senderId, err := uuid.Parse(req.SenderId)
	if err != nil {
		return &pb.SendFriendResponse{
			Status: http.StatusBadRequest,
			Error:  "user service: invalid sender id",
		}, nil
	}
	receiverId, err := uuid.Parse(req.ReceiverId)
	if err != nil {
		return &pb.SendFriendResponse{
			Status: http.StatusBadRequest,
			Error:  "user service: invalid receiver id",
		}, nil
	}

	res, err := s.service.SendFriend(senderId, receiverId)
	if err != nil {
		return &pb.SendFriendResponse{
			Status: http.StatusBadGateway,
			Error:  err.Error(),
		}, nil
	}

	return &pb.SendFriendResponse{
		Status: http.StatusOK,
		FriendRequest: &pb.FriendRequest{
			Id:         res.Id.String(),
			SenderId:   req.SenderId,
			ReceiverId: req.ReceiverId,
			IsAccepted: false,
		},
	}, nil
}

func (s UsersServer) AcceptFriend(_ context.Context, req *pb.AcceptDeclineFriendRequest) (*pb.DefaultResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return &pb.DefaultResponse{
			Status: http.StatusBadRequest,
			Error:  "user service: invalid user id",
		}, nil
	}
	requestId, err := uuid.Parse(req.RequestId)
	if err != nil {
		return &pb.DefaultResponse{
			Status: http.StatusBadRequest,
			Error:  "user service: invalid request id",
		}, nil
	}

	err = s.service.AcceptFriend(userId, requestId)
	if err != nil {
		return &pb.DefaultResponse{
			Status: http.StatusBadGateway,
			Error:  err.Error(),
		}, nil
	}

	return &pb.DefaultResponse{
		Status: http.StatusOK,
	}, nil
}

func (s UsersServer) DeclineFriend(_ context.Context, req *pb.AcceptDeclineFriendRequest) (*pb.DefaultResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return &pb.DefaultResponse{
			Status: http.StatusBadRequest,
			Error:  "user service: invalid user id",
		}, nil
	}
	requestId, err := uuid.Parse(req.RequestId)
	if err != nil {
		return &pb.DefaultResponse{
			Status: http.StatusBadRequest,
			Error:  "user service: invalid request id",
		}, nil
	}

	err = s.service.DeclineFriend(userId, requestId)
	if err != nil {
		return &pb.DefaultResponse{
			Status: http.StatusBadGateway,
			Error:  err.Error(),
		}, nil
	}

	return &pb.DefaultResponse{
		Status: http.StatusOK,
	}, nil
}

func (s UsersServer) Unfriend(_ context.Context, req *pb.UnfriendRequest) (*pb.DefaultResponse, error) {
	senderId, err := uuid.Parse(req.SenderId)
	if err != nil {
		return &pb.DefaultResponse{
			Status: http.StatusBadRequest,
			Error:  "user service: invalid sender id",
		}, nil
	}
	receiverId, err := uuid.Parse(req.ReceiverId)
	if err != nil {
		return &pb.DefaultResponse{
			Status: http.StatusBadRequest,
			Error:  "user service: invalid receiver id",
		}, nil
	}

	err = s.service.Unfriend(senderId, receiverId)
	if err != nil {
		return &pb.DefaultResponse{
			Status: http.StatusBadGateway,
			Error:  err.Error(),
		}, nil
	}

	return &pb.DefaultResponse{
		Status: http.StatusOK,
	}, nil
}
