package servers

import (
	"user-service/internal/pb"
	"user-service/internal/services"
)

type UsersServer struct {
	service *services.UsersService
	pb.UnimplementedUsersServiceServer
}
