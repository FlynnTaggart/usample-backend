package auth

import (
	"api-gateway-service/internal/auth/handlers"
	"api-gateway-service/pkg/zapadapter"
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(a *fiber.App, URL string, logger zapadapter.ZapAdapter) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(URL, logger),
	}

	routes := a.Group("/auth")
	routes.Post("/register", svc.Register)
	routes.Post("/login", svc.Login)

	return svc
}

func (svc *ServiceClient) Register(ctx *fiber.Ctx) error {
	return handlers.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *fiber.Ctx) error {
	return handlers.Login(ctx, svc.Client)
}
