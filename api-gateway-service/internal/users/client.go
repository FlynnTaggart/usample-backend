package users

import (
	"api-gateway-service/internal/pb/users_pb"
	"api-gateway-service/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client users_pb.UsersServiceClient
}

func InitServiceClient(serviceUrl string, log logger.Logger) users_pb.UsersServiceClient {
	cc, err := grpc.Dial(serviceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Error("Could not connect to users service client", map[string]interface{}{
			"error": err.Error(),
		})
	}

	return users_pb.NewUsersServiceClient(cc)
}
