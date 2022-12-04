package auth

import (
	"api-gateway-service/internal/pb/auth_pb"
	"api-gateway-service/pkg/zapadapter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client auth_pb.AuthServiceClient
}

func InitServiceClient(serviceUrl string, log zapadapter.ZapAdapter) auth_pb.AuthServiceClient {
	cc, err := grpc.Dial(serviceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Error("Could not connect to auth service client", map[string]interface{}{
			"error": err.Error(),
		})
	}

	return auth_pb.NewAuthServiceClient(cc)
}
