package users

import (
	"api-gateway-service/internal/pb/samples_pb"
	"api-gateway-service/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client samples_pb.SamplesServiceClient
}

func InitServiceClient(serviceUrl string, log logger.Logger) samples_pb.SamplesServiceClient {
	cc, err := grpc.Dial(serviceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Error("Could not connect to users service client", map[string]interface{}{
			"error": err.Error(),
		})
	}

	return samples_pb.NewSamplesServiceClient(cc)
}
