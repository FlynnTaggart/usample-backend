package servers

import (
	"auth-service/internal/pb"
	"auth-service/internal/services"
	"strings"

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

func (s AuthServer) Register(_ context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if len(req.Password) == 0 || len(req.Email) == 0 {
		return &pb.RegisterResponse{
			Status: http.StatusBadRequest,
			Error:  "auth server: register: empty email or password",
		}, nil
	}

	userId, err := s.service.Register(req.Email, req.Password)

	if err != nil && strings.Contains(err.Error(), "timeout") {
		return &pb.RegisterResponse{
			Status: http.StatusRequestTimeout,
			Error:  err.Error(),
		}, nil
	} else if err != nil {
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

func (s AuthServer) Login(_ context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if len(req.Password) == 0 || len(req.Email) == 0 {
		return &pb.LoginResponse{
			Status: http.StatusBadRequest,
			Error:  "auth server: register: empty email or password",
		}, nil
	}

	token, err := s.service.Login(req.Email, req.Password)

	if err != nil && strings.Contains(err.Error(), "timeout") {
		return &pb.LoginResponse{
			Status: http.StatusRequestTimeout,
			Error:  err.Error(),
		}, nil
	} else if err != nil {
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

func (s AuthServer) Validate(_ context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	id, err := s.service.Validate(req.Token)

	if err != nil && strings.Contains(err.Error(), "timeout") {
		return &pb.ValidateResponse{
			Status: http.StatusRequestTimeout,
			Error:  err.Error(),
		}, nil
	} else if err != nil {
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
