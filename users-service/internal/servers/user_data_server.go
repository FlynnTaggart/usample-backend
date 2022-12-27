package servers

import (
	"context"
	"github.com/google/uuid"
	"net/http"
	"user-service/internal/models"
	"user-service/internal/pb"
)

func (s UsersServer) CreateUser(_ context.Context, req *pb.User) (*pb.DefaultResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return &pb.DefaultResponse{
			Status: http.StatusBadRequest,
			Error:  "user service: invalid user id",
		}, nil
	}

	user := &models.User{
		Id:                id,
		Nickname:          req.Nickname,
		FirstName:         req.FirstName,
		SecondName:        req.SecondName,
		DefaultAccessType: models.SampleAccessType(req.DefaultAccessType),
		UserType:          models.UserType(req.UserType),
		Bio:               req.Bio,
	}

	err = s.service.CreateUser(user)

	if err != nil {
		return &pb.DefaultResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	return &pb.DefaultResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s UsersServer) GetUsers(_ context.Context, req *pb.GetUsersRequest) (*pb.UsersResponse, error) {
	limit, offset := req.Limit, req.Offset

	res, err := s.service.GetUsers(limit, offset)
	if err != nil {
		return &pb.UsersResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	resp := &pb.UsersResponse{
		Status: http.StatusOK,
		Users:  []*pb.User{},
	}

	for _, v := range res {
		resp.Users = append(resp.Users, &pb.User{
			Id:                v.Id.String(),
			Nickname:          v.Nickname,
			FirstName:         v.FirstName,
			SecondName:        v.SecondName,
			DefaultAccessType: pb.SampleAccessType(v.DefaultAccessType),
			UserType:          pb.UserType(v.UserType),
			Bio:               v.Bio,
		})
	}

	return resp, nil
}

func (s UsersServer) GetUser(_ context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return &pb.UserResponse{
			Status: http.StatusBadRequest,
			Error:  "user service: invalid user id",
		}, nil
	}
	res, err := s.service.GetUser(id)
	if err != nil {
		return &pb.UserResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	resp := &pb.User{
		Id:                res.Id.String(),
		Nickname:          res.Nickname,
		FirstName:         res.FirstName,
		SecondName:        res.SecondName,
		DefaultAccessType: pb.SampleAccessType(res.DefaultAccessType),
		UserType:          pb.UserType(res.UserType),
		Bio:               res.Bio,
	}

	return &pb.UserResponse{User: resp, Status: http.StatusOK}, nil
}

func (s UsersServer) GetUsersByNicknamePrefix(_ context.Context, req *pb.GetUsersByNicknamePrefixRequest) (*pb.UsersResponse, error) {
	prefix, limit, offset := req.Query, req.Limit, req.Offset

	res, err := s.service.GetUsersByNicknamePrefix(prefix, limit, offset)
	if err != nil {
		return &pb.UsersResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	resp := &pb.UsersResponse{
		Status: http.StatusOK,
		Users:  []*pb.User{},
	}

	for _, v := range res {
		resp.Users = append(resp.Users, &pb.User{
			Id:                v.Id.String(),
			Nickname:          v.Nickname,
			FirstName:         v.FirstName,
			SecondName:        v.SecondName,
			DefaultAccessType: pb.SampleAccessType(v.DefaultAccessType),
			UserType:          pb.UserType(v.UserType),
			Bio:               v.Bio,
		})
	}

	return resp, nil
}

func (s UsersServer) GetUserByNickname(_ context.Context, req *pb.GetUserByNicknameRequest) (*pb.UserResponse, error) {
	if len(req.Nickname) == 0 {
		return &pb.UserResponse{
			Status: http.StatusBadRequest,
			Error:  "gateway: get user by nickname: empty nickname",
		}, nil
	}
	res, err := s.service.GetUserByNickname(req.GetNickname())
	if err != nil {
		return &pb.UserResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	resp := &pb.User{
		Id:                res.Id.String(),
		Nickname:          res.Nickname,
		FirstName:         res.FirstName,
		SecondName:        res.SecondName,
		DefaultAccessType: pb.SampleAccessType(res.DefaultAccessType),
		UserType:          pb.UserType(res.UserType),
		Bio:               res.Bio,
	}

	return &pb.UserResponse{User: resp, Status: http.StatusOK}, nil
}

func (s UsersServer) UpdateUserInfo(_ context.Context, req *pb.User) (*pb.DefaultResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return &pb.DefaultResponse{
			Status: http.StatusBadRequest,
			Error:  "user service: invalid user id",
		}, nil
	}

	user := &models.User{
		Id:                id,
		Nickname:          req.Nickname,
		FirstName:         req.FirstName,
		SecondName:        req.SecondName,
		DefaultAccessType: models.SampleAccessType(req.DefaultAccessType),
		UserType:          models.UserType(req.UserType),
		Bio:               req.Bio,
	}

	err = s.service.UpdateUserInfo(user)
	if err != nil {
		return &pb.DefaultResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	return &pb.DefaultResponse{
		Status: http.StatusOK,
	}, nil
}
