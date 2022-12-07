package auth

import (
	"api-gateway-service/internal/auth/handlers"
	"api-gateway-service/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

func InitializeLoginRoute(a fiber.Router, URL string, logger logger.Logger) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(URL, logger),
	}
	a.Post("/login", func(ctx *fiber.Ctx) error {
		return handlers.Login(ctx, svc.Client)
	})

	return svc
}
