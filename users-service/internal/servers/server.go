package servers

import (
	"user-service/internal/pb"
	"user-service/internal/services"
	"user-service/pkg/logger"
)

type UsersServer struct {
	service *services.UsersService
	log     logger.Logger
	pb.UnimplementedUsersServiceServer
}

func NewUsersServer(service *services.UsersService, log logger.Logger) *UsersServer {
	return &UsersServer{service: service, log: log}
}
