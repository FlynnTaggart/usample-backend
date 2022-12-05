package auth

import (
	"api-gateway-service/internal/auth/handlers"
	"api-gateway-service/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(a fiber.Router, URL string, logger logger.Logger) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(URL, logger),
	}

	a.Post("/register", svc.Register)
	a.Post("/login", svc.Login)

	return svc
}

func (svc *ServiceClient) Register(ctx *fiber.Ctx) error {
	return handlers.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *fiber.Ctx) error {
	return handlers.Login(ctx, svc.Client)
}
