package servers

import (
	"auth-service/internal/pb"
	"auth-service/internal/services"

	"context"
	"net/http"
)

type AuthServer struct {
	service *services.AuthService
	pb.UnimplementedAuthServiceServer
}

func NewAuthServer(service *services.AuthService) *AuthServer {
	return &AuthServer{service: service}
}

func (s AuthServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	userId, err := s.service.Register(req.Email, req.Password)

	if err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
		UserId: userId,
	}, nil
}

func (s AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := s.service.Login(req.Email, req.Password)

	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (s AuthServer) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	id, err := s.service.Validate(req.Token)

	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusOK,
			Error:  err.Error(),
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: id.String(),
	}, nil
}
